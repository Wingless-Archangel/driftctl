// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/cloudskiff/driftctl/pkg/helpers"
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const AwsSnsTopicPolicyResourceType = "aws_sns_topic_policy"

type AwsSnsTopicPolicy struct {
	Arn    *string    `cty:"arn"`
	Id     string     `cty:"id" computed:"true"`
	Policy *string    `cty:"policy" jsonstring:"true"`
	CtyVal *cty.Value `diff:"-"`
}

func (r *AwsSnsTopicPolicy) TerraformId() string {
	return r.Id
}

func (r *AwsSnsTopicPolicy) TerraformType() string {
	return AwsSnsTopicPolicyResourceType
}

func (r *AwsSnsTopicPolicy) CtyValue() *cty.Value {
	return r.CtyVal
}

func initSnsTopicPolicyMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.UpdateSchema(AwsSnsTopicPolicyResourceType, map[string]func(attributeSchema *resource.AttributeSchema){
		"policy": func(attributeSchema *resource.AttributeSchema) {
			attributeSchema.JsonString = true
		},
	})

	resourceSchemaRepository.SetNormalizeFunc(AwsSnsTopicPolicyResourceType, func(val *resource.Attributes) {
		jsonString, err := helpers.NormalizeJsonString((*val)["policy"])
		if err != nil {
			return
		}
		val.SafeSet([]string{"policy"}, jsonString)
	})
}
