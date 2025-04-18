/* SPDX-License-Identifier: (GPL-2.0-only OR BSD-2-Clause) */
/* Copyright Authors of Cilium */

#pragma once

#include "common.h"
#include "maps.h"
#include "ipv6.h"
#include "ipv4.h"
#include "eps.h"
#include "eth.h"
#include "dbg.h"
#include "l4.h"
#include "icmp6.h"
#include "csum.h"
#include "token_bucket.h"

#ifdef ENABLE_IPV6
static __always_inline int ipv6_l3(struct __ctx_buff *ctx, int l3_off,
				   const __u8 *smac, const __u8 *dmac,
				   __u8 __maybe_unused direction)
{
	int ret;

	ret = ipv6_dec_hoplimit(ctx, l3_off);
	if (IS_ERR(ret)) {
#ifndef SKIP_ICMPV6_HOPLIMIT_HANDLING
		if (ret == DROP_TTL_EXCEEDED)
			return icmp6_send_time_exceeded(ctx, l3_off, direction);
#endif
		return ret;
	}

	if (smac && eth_store_saddr(ctx, smac, 0) < 0)
		return DROP_WRITE_ERROR;
	if (dmac && eth_store_daddr(ctx, dmac, 0) < 0)
		return DROP_WRITE_ERROR;

	return CTX_ACT_OK;
}
#endif /* ENABLE_IPV6 */

static __always_inline int ipv4_l3(struct __ctx_buff *ctx, int l3_off,
				   const __u8 *smac, const __u8 *dmac,
				   struct iphdr *ip4)
{
	int ret;

	ret = ipv4_dec_ttl(ctx, l3_off, ip4);
	/* FIXME: Send ICMP TTL */
	if (IS_ERR(ret))
		return ret;

	if (smac && eth_store_saddr(ctx, smac, 0) < 0)
		return DROP_WRITE_ERROR;
	if (dmac && eth_store_daddr(ctx, dmac, 0) < 0)
		return DROP_WRITE_ERROR;

	return CTX_ACT_OK;
}

/* Defines the calling convention for bpf_lxc's ingress policy tail-call.
 * Note that skb->tc_index is also passed through.
 *
 * As the callers (from-overlay, from-netdev, ...) are re-generated independently
 * from the policy tail-call of the inidividual endpoints, any change to this code
 * needs to be introduced with compatibility in mind.
 */
static __always_inline void
local_delivery_fill_meta(struct __ctx_buff *ctx, __u32 seclabel,
			 bool delivery_redirect, bool from_host,
			 bool from_tunnel, __u32 cluster_id)
{
	ctx_store_meta(ctx, CB_SRC_LABEL, seclabel);
	ctx_store_meta(ctx, CB_DELIVERY_REDIRECT, delivery_redirect ? 1 : 0);
	ctx_store_meta(ctx, CB_FROM_HOST, from_host ? 1 : 0);
	ctx_store_meta(ctx, CB_FROM_TUNNEL, from_tunnel ? 1 : 0);
	ctx_store_meta(ctx, CB_CLUSTER_ID_INGRESS, cluster_id);
}

#ifndef SKIP_POLICY_MAP
static __always_inline int
l3_local_delivery(struct __ctx_buff *ctx, __u32 seclabel,
		  __u32 magic __maybe_unused,
		  const struct endpoint_info *ep __maybe_unused,
		  __u8 direction __maybe_unused,
		  bool from_host __maybe_unused,
		  bool from_tunnel __maybe_unused, __u32 cluster_id __maybe_unused)
{
#ifdef LOCAL_DELIVERY_METRICS
	/*
	 * Special LXC case for updating egress forwarding metrics.
	 * Note that the packet could still be dropped but it would show up
	 * as an ingress drop counter in metrics.
	 */
	update_metrics(ctx_full_len(ctx), direction, REASON_FORWARDED);
#endif

	if (direction == METRIC_INGRESS && !from_host) {
		/*
		 * Traffic from nodes, local endpoints, or hairpin connections is ignored
		 */
		int ret;

		ret = accept(ctx, ep->lxc_id);
		if (IS_ERR(ret))
			return ret;
	}

/*
 * When BPF host routing is enabled we need to check policies at source, as in
 * this case the skb is delivered directly to pod's namespace and the ingress
 * policy (the cil_to_container BPF program) is bypassed.
 */
#if defined(USE_BPF_PROG_FOR_INGRESS_POLICY) && \
    !defined(ENABLE_HOST_ROUTING)
	set_identity_mark(ctx, seclabel, magic);

# if !defined(ENABLE_NODEPORT)
	/* In tunneling mode, we execute this code to send the packet from
	 * cilium_vxlan to lxc*. If we're using kube-proxy, we don't want to use
	 * redirect() because that would bypass conntrack and the reverse DNAT.
	 * Thus, we send packets to the stack, but since they have the wrong
	 * Ethernet addresses, we need to mark them as PACKET_HOST or the kernel
	 * will drop them.
	 */
	if (from_tunnel) {
		ctx_change_type(ctx, PACKET_HOST);
		return CTX_ACT_OK;
	}
# endif /* !ENABLE_NODEPORT */

	return redirect_ep(ctx, ep->ifindex, from_host, from_tunnel);
#else

	/* Jumps to destination pod's BPF program to enforce ingress policies. */
	local_delivery_fill_meta(ctx, seclabel, true, from_host, from_tunnel, cluster_id);
	return tail_call_policy(ctx, ep->lxc_id);
#endif
}

#ifdef ENABLE_IPV6
/* Performs IPv6 L2/L3 handling and delivers the packet to the destination pod
 * on the same node, either via the stack or via a redirect call.
 * Depending on the configuration, it may also enforce ingress policies for the
 * destination pod via a tail call.
 */
static __always_inline int ipv6_local_delivery(struct __ctx_buff *ctx, int l3_off,
					       __u32 seclabel, __u32 magic,
					       const struct endpoint_info *ep,
					       __u8 direction, bool from_host,
					       bool from_tunnel)
{
	mac_t router_mac = ep->node_mac;
	mac_t lxc_mac = ep->mac;
	int ret;

	cilium_dbg(ctx, DBG_LOCAL_DELIVERY, ep->lxc_id, seclabel);

	ret = ipv6_l3(ctx, l3_off, (__u8 *)&router_mac, (__u8 *)&lxc_mac, direction);
	if (ret != CTX_ACT_OK)
		return ret;

	return l3_local_delivery(ctx, seclabel, magic, ep, direction, from_host,
				 from_tunnel, 0);
}
#endif /* ENABLE_IPV6 */

/* Performs IPv4 L2/L3 handling and delivers the packet to the destination pod
 * on the same node, either via the stack or via a redirect call.
 * Depending on the configuration, it may also enforce ingress policies for the
 * destination pod via a tail call.
 */
static __always_inline int ipv4_local_delivery(struct __ctx_buff *ctx, int l3_off,
					       __u32 seclabel, __u32 magic,
					       struct iphdr *ip4,
					       const struct endpoint_info *ep,
					       __u8 direction, bool from_host,
					       bool from_tunnel, __u32 cluster_id)
{
	mac_t router_mac = ep->node_mac;
	mac_t lxc_mac = ep->mac;
	int ret;

	cilium_dbg(ctx, DBG_LOCAL_DELIVERY, ep->lxc_id, seclabel);

	ret = ipv4_l3(ctx, l3_off, (__u8 *) &router_mac, (__u8 *) &lxc_mac, ip4);
	if (ret != CTX_ACT_OK)
		return ret;

	return l3_local_delivery(ctx, seclabel, magic, ep, direction, from_host,
				 from_tunnel, cluster_id);
}
#endif /* SKIP_POLICY_MAP */
