package cmd

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/spf13/viper"
)

func createSSMService() *ssm.SSM {
	endpoint := viper.GetString("endpoint-url")

	options := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}

	if endpoint != "" {
		options.Config.MergeIn(aws.NewConfig().WithEndpoint(endpoint))
	}

	// Get config from ~/.aws/config file
	sess := session.Must(session.NewSessionWithOptions(options))

	// Create SSM service client
	return ssm.New(sess)
}
