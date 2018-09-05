package cmd

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/spf13/viper"
)

func createSSMService() *ssm.SSM {
	endpoint := viper.GetString("endpoint-url")

	config := aws.NewConfig().WithRegion("eu-west-1")

	if endpoint != "" {
		config = config.WithEndpoint(endpoint)
	}

	sess := session.Must(session.NewSession(config))

	// Create SSM service client
	return ssm.New(sess)
}
