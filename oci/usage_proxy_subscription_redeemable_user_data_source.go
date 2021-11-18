// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v52/usage"
)

func init() {
	RegisterDatasource("oci_usage_proxy_subscription_redeemable_user", UsageProxySubscriptionRedeemableUserDataSource())
}

func UsageProxySubscriptionRedeemableUserDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["subscription_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["tenancy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(UsageProxySubscriptionRedeemableUserResource(), fieldMap, readSingularUsageProxySubscriptionRedeemableUser)
}

func readSingularUsageProxySubscriptionRedeemableUser(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionRedeemableUserDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).rewardsClient()

	return ReadResource(sync)
}

type UsageProxySubscriptionRedeemableUserDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.RewardsClient
	Res    *oci_usage_proxy.ListRedeemableUsersResponse
}

func (s *UsageProxySubscriptionRedeemableUserDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxySubscriptionRedeemableUserDataSourceCrud) Get() error {
	request := oci_usage_proxy.ListRedeemableUsersRequest{}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "usage_proxy")

	response, err := s.Client.ListRedeemableUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *UsageProxySubscriptionRedeemableUserDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("UsageProxySubscriptionRedeemableUserDataSource-", UsageProxySubscriptionRedeemableUserDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RedeemableUserSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
