package dao

import (
	"context"
	"encoding/json"
	"github.com/dashenmiren/EdgeCommon/pkg/errors"
	"github.com/dashenmiren/EdgeCommon/pkg/rpc/pb"
	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs/ipconfigs"
)

var SharedIPListDAO = new(IPListDAO)

type IPListDAO struct {
	BaseDAO
}

// FindAllowIPListIdWithServerId 查找网站的允许IP列表
func (this *IPListDAO) FindAllowIPListIdWithServerId(ctx context.Context, serverId int64) (int64, error) {
	webConfig, err := SharedHTTPWebDAO.FindWebConfigWithServerId(ctx, serverId)
	if err != nil {
		return 0, err
	}
	if webConfig == nil {
		return 0, nil
	}
	if webConfig.FirewallPolicy == nil || webConfig.FirewallPolicy.Inbound == nil || webConfig.FirewallPolicy.Inbound.AllowListRef == nil {
		return 0, nil
	}
	return webConfig.FirewallPolicy.Inbound.AllowListRef.ListId, nil
}

// FindDenyIPListIdWithServerId 查找网站的禁止IP列表
func (this *IPListDAO) FindDenyIPListIdWithServerId(ctx context.Context, serverId int64) (int64, error) {
	webConfig, err := SharedHTTPWebDAO.FindWebConfigWithServerId(ctx, serverId)
	if err != nil {
		return 0, err
	}
	if webConfig == nil {
		return 0, nil
	}
	if webConfig.FirewallPolicy == nil || webConfig.FirewallPolicy.Inbound == nil || webConfig.FirewallPolicy.Inbound.DenyListRef == nil {
		return 0, nil
	}
	return webConfig.FirewallPolicy.Inbound.DenyListRef.ListId, nil
}

// FindGreyIPListIdWithServerId 查找网站的IP灰名单
func (this *IPListDAO) FindGreyIPListIdWithServerId(ctx context.Context, serverId int64) (int64, error) {
	webConfig, err := SharedHTTPWebDAO.FindWebConfigWithServerId(ctx, serverId)
	if err != nil {
		return 0, err
	}
	if webConfig == nil {
		return 0, nil
	}
	if webConfig.FirewallPolicy == nil || webConfig.FirewallPolicy.Inbound == nil || webConfig.FirewallPolicy.Inbound.GreyListRef == nil {
		return 0, nil
	}
	return webConfig.FirewallPolicy.Inbound.GreyListRef.ListId, nil
}

// CreateIPListForServerId 为服务创建IP名单
func (this *IPListDAO) CreateIPListForServerId(ctx context.Context, serverId int64, listType string) (int64, error) {
	webConfig, err := SharedHTTPWebDAO.FindWebConfigWithServerId(ctx, serverId)
	if err != nil {
		return 0, err
	}
	if webConfig == nil {
		return 0, nil
	}
	if webConfig.FirewallPolicy == nil || webConfig.FirewallPolicy.Id == 0 {
		isOn := webConfig.FirewallRef != nil && webConfig.FirewallRef.IsOn
		_, err = SharedHTTPWebDAO.InitEmptyHTTPFirewallPolicy(ctx, 0, serverId, webConfig.Id, isOn)
		if err != nil {
			return 0, errors.Wrap(err)
		}
		webConfig, err = SharedHTTPWebDAO.FindWebConfigWithServerId(ctx, serverId)
		if err != nil {
			return 0, err
		}
		if webConfig == nil {
			return 0, nil
		}
		if webConfig.FirewallPolicy == nil {
			return 0, nil
		}
	}

	var inbound = webConfig.FirewallPolicy.Inbound
	if inbound == nil {
		inbound = &firewallconfigs.HTTPFirewallInboundConfig{
			IsOn: true,
		}
	}
	if listType == ipconfigs.IPListTypeWhite {
		if inbound.AllowListRef == nil {
			inbound.AllowListRef = &ipconfigs.IPListRef{
				IsOn: true,
			}
		}
		if inbound.AllowListRef.ListId > 0 {
			return inbound.AllowListRef.ListId, nil
		}
	} else if listType == ipconfigs.IPListTypeBlack {
		if inbound.DenyListRef == nil {
			inbound.DenyListRef = &ipconfigs.IPListRef{
				IsOn: true,
			}
		}
		if inbound.DenyListRef.ListId > 0 {
			return inbound.DenyListRef.ListId, nil
		}
	} else if listType == ipconfigs.IPListTypeGrey {
		if inbound.GreyListRef == nil {
			inbound.GreyListRef = &ipconfigs.IPListRef{
				IsOn: true,
			}
		}
		if inbound.GreyListRef.ListId > 0 {
			return inbound.DenyListRef.ListId, nil
		}
	}

	ipListResp, err := this.RPC().IPListRPC().CreateIPList(ctx, &pb.CreateIPListRequest{
		Type:        listType,
		Name:        "IP名单",
		Code:        listType,
		ServerId:    serverId,
		TimeoutJSON: nil,
	})
	if err != nil {
		return 0, errors.Wrap(err)
	}

	if listType == ipconfigs.IPListTypeWhite {
		inbound.AllowListRef.ListId = ipListResp.IpListId
	} else if listType == ipconfigs.IPListTypeBlack {
		inbound.DenyListRef.ListId = ipListResp.IpListId
	} else if listType == ipconfigs.IPListTypeGrey {
		inbound.GreyListRef.ListId = ipListResp.IpListId
	}
	inboundJSON, err := json.Marshal(inbound)
	if err != nil {
		return 0, errors.Wrap(err)
	}
	_, err = this.RPC().HTTPFirewallPolicyRPC().UpdateHTTPFirewallInboundConfig(ctx, &pb.UpdateHTTPFirewallInboundConfigRequest{
		HttpFirewallPolicyId: webConfig.FirewallPolicy.Id,
		InboundJSON:          inboundJSON,
	})
	if err != nil {
		return 0, errors.Wrap(err)
	}

	return ipListResp.IpListId, nil
}
