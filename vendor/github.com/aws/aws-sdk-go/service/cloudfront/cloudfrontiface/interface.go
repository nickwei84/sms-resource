// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudfrontiface provides an interface for the Amazon CloudFront.
package cloudfrontiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

// CloudFrontAPI is the interface type for cloudfront.CloudFront.
type CloudFrontAPI interface {
	CreateCloudFrontOriginAccessIdentityRequest(*cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.CreateCloudFrontOriginAccessIdentityOutput)

	CreateCloudFrontOriginAccessIdentity(*cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error)

	CreateDistributionRequest(*cloudfront.CreateDistributionInput) (*request.Request, *cloudfront.CreateDistributionOutput)

	CreateDistribution(*cloudfront.CreateDistributionInput) (*cloudfront.CreateDistributionOutput, error)

	CreateInvalidationRequest(*cloudfront.CreateInvalidationInput) (*request.Request, *cloudfront.CreateInvalidationOutput)

	CreateInvalidation(*cloudfront.CreateInvalidationInput) (*cloudfront.CreateInvalidationOutput, error)

	CreateStreamingDistributionRequest(*cloudfront.CreateStreamingDistributionInput) (*request.Request, *cloudfront.CreateStreamingDistributionOutput)

	CreateStreamingDistribution(*cloudfront.CreateStreamingDistributionInput) (*cloudfront.CreateStreamingDistributionOutput, error)

	DeleteCloudFrontOriginAccessIdentityRequest(*cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.DeleteCloudFrontOriginAccessIdentityOutput)

	DeleteCloudFrontOriginAccessIdentity(*cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error)

	DeleteDistributionRequest(*cloudfront.DeleteDistributionInput) (*request.Request, *cloudfront.DeleteDistributionOutput)

	DeleteDistribution(*cloudfront.DeleteDistributionInput) (*cloudfront.DeleteDistributionOutput, error)

	DeleteStreamingDistributionRequest(*cloudfront.DeleteStreamingDistributionInput) (*request.Request, *cloudfront.DeleteStreamingDistributionOutput)

	DeleteStreamingDistribution(*cloudfront.DeleteStreamingDistributionInput) (*cloudfront.DeleteStreamingDistributionOutput, error)

	GetCloudFrontOriginAccessIdentityRequest(*cloudfront.GetCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.GetCloudFrontOriginAccessIdentityOutput)

	GetCloudFrontOriginAccessIdentity(*cloudfront.GetCloudFrontOriginAccessIdentityInput) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error)

	GetCloudFrontOriginAccessIdentityConfigRequest(*cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*request.Request, *cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput)

	GetCloudFrontOriginAccessIdentityConfig(*cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error)

	GetDistributionRequest(*cloudfront.GetDistributionInput) (*request.Request, *cloudfront.GetDistributionOutput)

	GetDistribution(*cloudfront.GetDistributionInput) (*cloudfront.GetDistributionOutput, error)

	GetDistributionConfigRequest(*cloudfront.GetDistributionConfigInput) (*request.Request, *cloudfront.GetDistributionConfigOutput)

	GetDistributionConfig(*cloudfront.GetDistributionConfigInput) (*cloudfront.GetDistributionConfigOutput, error)

	GetInvalidationRequest(*cloudfront.GetInvalidationInput) (*request.Request, *cloudfront.GetInvalidationOutput)

	GetInvalidation(*cloudfront.GetInvalidationInput) (*cloudfront.GetInvalidationOutput, error)

	GetStreamingDistributionRequest(*cloudfront.GetStreamingDistributionInput) (*request.Request, *cloudfront.GetStreamingDistributionOutput)

	GetStreamingDistribution(*cloudfront.GetStreamingDistributionInput) (*cloudfront.GetStreamingDistributionOutput, error)

	GetStreamingDistributionConfigRequest(*cloudfront.GetStreamingDistributionConfigInput) (*request.Request, *cloudfront.GetStreamingDistributionConfigOutput)

	GetStreamingDistributionConfig(*cloudfront.GetStreamingDistributionConfigInput) (*cloudfront.GetStreamingDistributionConfigOutput, error)

	ListCloudFrontOriginAccessIdentitiesRequest(*cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*request.Request, *cloudfront.ListCloudFrontOriginAccessIdentitiesOutput)

	ListCloudFrontOriginAccessIdentities(*cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error)

	ListCloudFrontOriginAccessIdentitiesPages(*cloudfront.ListCloudFrontOriginAccessIdentitiesInput, func(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, bool) bool) error

	ListDistributionsRequest(*cloudfront.ListDistributionsInput) (*request.Request, *cloudfront.ListDistributionsOutput)

	ListDistributions(*cloudfront.ListDistributionsInput) (*cloudfront.ListDistributionsOutput, error)

	ListDistributionsPages(*cloudfront.ListDistributionsInput, func(*cloudfront.ListDistributionsOutput, bool) bool) error

	ListDistributionsByWebACLIdRequest(*cloudfront.ListDistributionsByWebACLIdInput) (*request.Request, *cloudfront.ListDistributionsByWebACLIdOutput)

	ListDistributionsByWebACLId(*cloudfront.ListDistributionsByWebACLIdInput) (*cloudfront.ListDistributionsByWebACLIdOutput, error)

	ListInvalidationsRequest(*cloudfront.ListInvalidationsInput) (*request.Request, *cloudfront.ListInvalidationsOutput)

	ListInvalidations(*cloudfront.ListInvalidationsInput) (*cloudfront.ListInvalidationsOutput, error)

	ListInvalidationsPages(*cloudfront.ListInvalidationsInput, func(*cloudfront.ListInvalidationsOutput, bool) bool) error

	ListStreamingDistributionsRequest(*cloudfront.ListStreamingDistributionsInput) (*request.Request, *cloudfront.ListStreamingDistributionsOutput)

	ListStreamingDistributions(*cloudfront.ListStreamingDistributionsInput) (*cloudfront.ListStreamingDistributionsOutput, error)

	ListStreamingDistributionsPages(*cloudfront.ListStreamingDistributionsInput, func(*cloudfront.ListStreamingDistributionsOutput, bool) bool) error

	UpdateCloudFrontOriginAccessIdentityRequest(*cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.UpdateCloudFrontOriginAccessIdentityOutput)

	UpdateCloudFrontOriginAccessIdentity(*cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error)

	UpdateDistributionRequest(*cloudfront.UpdateDistributionInput) (*request.Request, *cloudfront.UpdateDistributionOutput)

	UpdateDistribution(*cloudfront.UpdateDistributionInput) (*cloudfront.UpdateDistributionOutput, error)

	UpdateStreamingDistributionRequest(*cloudfront.UpdateStreamingDistributionInput) (*request.Request, *cloudfront.UpdateStreamingDistributionOutput)

	UpdateStreamingDistribution(*cloudfront.UpdateStreamingDistributionInput) (*cloudfront.UpdateStreamingDistributionOutput, error)
}

var _ CloudFrontAPI = (*cloudfront.CloudFront)(nil)
