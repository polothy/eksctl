package cloudformation

import (
	"encoding/json"
)

// AWSEC2NetworkInterface_PrivateIpAddressSpecification AWS CloudFormation Resource (AWS::EC2::NetworkInterface.PrivateIpAddressSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html
type AWSEC2NetworkInterface_PrivateIpAddressSpecification struct {

	// Primary AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html#cfn-ec2-networkinterface-privateipspecification-primary
	Primary *Value `json:"Primary,omitempty"`

	// PrivateIpAddress AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html#cfn-ec2-networkinterface-privateipspecification-privateipaddress
	PrivateIpAddress *Value `json:"PrivateIpAddress,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkInterface_PrivateIpAddressSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkInterface.PrivateIpAddressSpecification"
}

func (r *AWSEC2NetworkInterface_PrivateIpAddressSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
