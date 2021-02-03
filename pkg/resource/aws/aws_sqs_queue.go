// GENERATED, DO NOT EDIT THIS FILE
package aws

const AwsSqsQueueResourceType = "aws_sqs_queue"

type AwsSqsQueue struct {
	Arn                          *string           `cty:"arn" computed:"true"`
	ContentBasedDeduplication    *bool             `cty:"content_based_deduplication"`
	DelaySeconds                 *int              `cty:"delay_seconds"`
	FifoQueue                    *bool             `cty:"fifo_queue"`
	Id                           string            `cty:"id" computed:"true"`
	KmsDataKeyReusePeriodSeconds *int              `cty:"kms_data_key_reuse_period_seconds" computed:"true"`
	KmsMasterKeyId               *string           `cty:"kms_master_key_id"`
	MaxMessageSize               *int              `cty:"max_message_size"`
	MessageRetentionSeconds      *int              `cty:"message_retention_seconds"`
	Name                         *string           `cty:"name" computed:"true"`
	NamePrefix                   *string           `cty:"name_prefix"`
	Policy                       *string           `cty:"policy" computed:"true"`
	ReceiveWaitTimeSeconds       *int              `cty:"receive_wait_time_seconds"`
	RedrivePolicy                *string           `cty:"redrive_policy"`
	Tags                         map[string]string `cty:"tags"`
	VisibilityTimeoutSeconds     *int              `cty:"visibility_timeout_seconds"`
}

func (r *AwsSqsQueue) TerraformId() string {
	return r.Id
}

func (r *AwsSqsQueue) TerraformType() string {
	return AwsSqsQueueResourceType
}
