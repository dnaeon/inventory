// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package tasks

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gardener/inventory/pkg/metrics"
)

var (
	// regionsDesc is the descriptor for a metric, which tracks the number
	// of collected AWS regions.
	regionsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metrics.Namespace, "", "aws_regions"),
		"A gauge which tracks the number of collected AWS Regions",
		[]string{"account_id"},
		nil,
	)

	// bucketsDesc is the descriptor for a metric, which tracks the number
	// of collected AWS S3 buckets.
	bucketsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metrics.Namespace, "", "aws_buckets"),
		"A gauge which tracks the number of collected AWS S3 buckets",
		[]string{"account_id", "region"},
		nil,
	)

	// imagesDesc is the descriptor for a metric, which tracks the number
	// of collected AWS AMI images.
	imagesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metrics.Namespace, "", "aws_images"),
		"A gauge which tracks the number of collected AWS AMI images",
		[]string{"account_id", "region"},
		nil,
	)

	// zonesDesc is the descriptor for a metric, which tracks the number
	// of collected AWS Availability Zones.
	zonesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metrics.Namespace, "", "aws_zones"),
		"A gauge which tracks the number of collected AWS AZs",
		[]string{"account_id", "region"},
		nil,
	)

	// vpcsDesc is the descriptor for a metric, which tracks the number
	// of collected AWS VPCs.
	vpcsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metrics.Namespace, "", "aws_vpcs"),
		"A gauge which tracks the number of collected AWS VPCs",
		[]string{"account_id", "region"},
		nil,
	)

	// subnetsDesc is the descriptor for a metric, which tracks the number
	// of collected AWS Subnets.
	subnetsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metrics.Namespace, "", "aws_subnets"),
		"A gauge which tracks the number of collected AWS Subnets",
		[]string{"account_id", "region", "vpc_id"},
		nil,
	)

	// instancesDesc is the descriptor for a metric, which tracks the number
	// of collected AWS EC2 instances.
	instancesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metrics.Namespace, "", "aws_instances"),
		"A gauge which tracks the number of collected AWS EC2 Instances",
		[]string{"account_id", "region", "vpc_id"},
		nil,
	)

	// loadBalancersDesc is the descriptor for a metric, which tracks the number
	// of collected AWS Elastic Load Balancers (ELBs).
	loadBalancersDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metrics.Namespace, "", "aws_load_balancers"),
		"A gauge which tracks the number of collected AWS ELBs",
		[]string{"account_id", "region", "vpc_id"},
		nil,
	)

	// netInterfacesDesc is the descriptor for a metric, which tracks the
	// number of collected AWS Elastic Network Interfaces (ENIs).
	netInterfacesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(metrics.Namespace, "", "aws_net_interfaces"),
		"A gauge which tracks the number of collected AWS ENIs",
		[]string{"account_id", "region", "vpc_id"},
		nil,
	)
)

// init registers the metrics with the [metrics.DefaultCollector]
func init() {
	metrics.DefaultCollector.AddDesc(
		regionsDesc,
		bucketsDesc,
		imagesDesc,
		zonesDesc,
		vpcsDesc,
		subnetsDesc,
		instancesDesc,
		loadBalancersDesc,
		netInterfacesDesc,
	)
}
