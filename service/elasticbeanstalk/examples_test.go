package elasticbeanstalk_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/datacratic/aws-sdk-go/aws"
	"github.com/datacratic/aws-sdk-go/aws/awsutil"
	"github.com/datacratic/aws-sdk-go/service/elasticbeanstalk"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleElasticBeanstalk_CheckDNSAvailability() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.CheckDNSAvailabilityInput{
		CNAMEPrefix: aws.String("DNSCnamePrefix"), // Required
	}
	resp, err := svc.CheckDNSAvailability(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_CreateApplication() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.CreateApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		Description:     aws.String("Description"),
	}
	resp, err := svc.CreateApplication(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_CreateApplicationVersion() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.CreateApplicationVersionInput{
		ApplicationName:       aws.String("ApplicationName"), // Required
		VersionLabel:          aws.String("VersionLabel"),    // Required
		AutoCreateApplication: aws.Boolean(true),
		Description:           aws.String("Description"),
		SourceBundle: &elasticbeanstalk.S3Location{
			S3Bucket: aws.String("S3Bucket"),
			S3Key:    aws.String("S3Key"),
		},
	}
	resp, err := svc.CreateApplicationVersion(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_CreateConfigurationTemplate() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.CreateConfigurationTemplateInput{
		ApplicationName: aws.String("ApplicationName"),           // Required
		TemplateName:    aws.String("ConfigurationTemplateName"), // Required
		Description:     aws.String("Description"),
		EnvironmentID:   aws.String("EnvironmentId"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			&elasticbeanstalk.ConfigurationOptionSetting{ // Required
				Namespace:  aws.String("OptionNamespace"),
				OptionName: aws.String("ConfigurationOptionName"),
				Value:      aws.String("ConfigurationOptionValue"),
			},
			// More values...
		},
		SolutionStackName: aws.String("SolutionStackName"),
		SourceConfiguration: &elasticbeanstalk.SourceConfiguration{
			ApplicationName: aws.String("ApplicationName"),
			TemplateName:    aws.String("ConfigurationTemplateName"),
		},
	}
	resp, err := svc.CreateConfigurationTemplate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_CreateEnvironment() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.CreateEnvironmentInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		EnvironmentName: aws.String("EnvironmentName"), // Required
		CNAMEPrefix:     aws.String("DNSCnamePrefix"),
		Description:     aws.String("Description"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			&elasticbeanstalk.ConfigurationOptionSetting{ // Required
				Namespace:  aws.String("OptionNamespace"),
				OptionName: aws.String("ConfigurationOptionName"),
				Value:      aws.String("ConfigurationOptionValue"),
			},
			// More values...
		},
		OptionsToRemove: []*elasticbeanstalk.OptionSpecification{
			&elasticbeanstalk.OptionSpecification{ // Required
				Namespace:  aws.String("OptionNamespace"),
				OptionName: aws.String("ConfigurationOptionName"),
			},
			// More values...
		},
		SolutionStackName: aws.String("SolutionStackName"),
		Tags: []*elasticbeanstalk.Tag{
			&elasticbeanstalk.Tag{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
		TemplateName: aws.String("ConfigurationTemplateName"),
		Tier: &elasticbeanstalk.EnvironmentTier{
			Name:    aws.String("String"),
			Type:    aws.String("String"),
			Version: aws.String("String"),
		},
		VersionLabel: aws.String("VersionLabel"),
	}
	resp, err := svc.CreateEnvironment(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_CreateStorageLocation() {
	svc := elasticbeanstalk.New(nil)

	var params *elasticbeanstalk.CreateStorageLocationInput
	resp, err := svc.CreateStorageLocation(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DeleteApplication() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DeleteApplicationInput{
		ApplicationName:     aws.String("ApplicationName"), // Required
		TerminateEnvByForce: aws.Boolean(true),
	}
	resp, err := svc.DeleteApplication(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DeleteApplicationVersion() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DeleteApplicationVersionInput{
		ApplicationName:    aws.String("ApplicationName"), // Required
		VersionLabel:       aws.String("VersionLabel"),    // Required
		DeleteSourceBundle: aws.Boolean(true),
	}
	resp, err := svc.DeleteApplicationVersion(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DeleteConfigurationTemplate() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DeleteConfigurationTemplateInput{
		ApplicationName: aws.String("ApplicationName"),           // Required
		TemplateName:    aws.String("ConfigurationTemplateName"), // Required
	}
	resp, err := svc.DeleteConfigurationTemplate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DeleteEnvironmentConfiguration() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DeleteEnvironmentConfigurationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		EnvironmentName: aws.String("EnvironmentName"), // Required
	}
	resp, err := svc.DeleteEnvironmentConfiguration(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DescribeApplicationVersions() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DescribeApplicationVersionsInput{
		ApplicationName: aws.String("ApplicationName"),
		VersionLabels: []*string{
			aws.String("VersionLabel"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeApplicationVersions(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DescribeApplications() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DescribeApplicationsInput{
		ApplicationNames: []*string{
			aws.String("ApplicationName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeApplications(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DescribeConfigurationOptions() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DescribeConfigurationOptionsInput{
		ApplicationName: aws.String("ApplicationName"),
		EnvironmentName: aws.String("EnvironmentName"),
		Options: []*elasticbeanstalk.OptionSpecification{
			&elasticbeanstalk.OptionSpecification{ // Required
				Namespace:  aws.String("OptionNamespace"),
				OptionName: aws.String("ConfigurationOptionName"),
			},
			// More values...
		},
		SolutionStackName: aws.String("SolutionStackName"),
		TemplateName:      aws.String("ConfigurationTemplateName"),
	}
	resp, err := svc.DescribeConfigurationOptions(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DescribeConfigurationSettings() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DescribeConfigurationSettingsInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		EnvironmentName: aws.String("EnvironmentName"),
		TemplateName:    aws.String("ConfigurationTemplateName"),
	}
	resp, err := svc.DescribeConfigurationSettings(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DescribeEnvironmentResources() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DescribeEnvironmentResourcesInput{
		EnvironmentID:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.DescribeEnvironmentResources(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DescribeEnvironments() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DescribeEnvironmentsInput{
		ApplicationName: aws.String("ApplicationName"),
		EnvironmentIDs: []*string{
			aws.String("EnvironmentId"), // Required
			// More values...
		},
		EnvironmentNames: []*string{
			aws.String("EnvironmentName"), // Required
			// More values...
		},
		IncludeDeleted:        aws.Boolean(true),
		IncludedDeletedBackTo: aws.Time(time.Now()),
		VersionLabel:          aws.String("VersionLabel"),
	}
	resp, err := svc.DescribeEnvironments(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_DescribeEvents() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.DescribeEventsInput{
		ApplicationName: aws.String("ApplicationName"),
		EndTime:         aws.Time(time.Now()),
		EnvironmentID:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
		MaxRecords:      aws.Long(1),
		NextToken:       aws.String("Token"),
		RequestID:       aws.String("RequestId"),
		Severity:        aws.String("EventSeverity"),
		StartTime:       aws.Time(time.Now()),
		TemplateName:    aws.String("ConfigurationTemplateName"),
		VersionLabel:    aws.String("VersionLabel"),
	}
	resp, err := svc.DescribeEvents(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_ListAvailableSolutionStacks() {
	svc := elasticbeanstalk.New(nil)

	var params *elasticbeanstalk.ListAvailableSolutionStacksInput
	resp, err := svc.ListAvailableSolutionStacks(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_RebuildEnvironment() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.RebuildEnvironmentInput{
		EnvironmentID:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.RebuildEnvironment(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_RequestEnvironmentInfo() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.RequestEnvironmentInfoInput{
		InfoType:        aws.String("EnvironmentInfoType"), // Required
		EnvironmentID:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.RequestEnvironmentInfo(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_RestartAppServer() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.RestartAppServerInput{
		EnvironmentID:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.RestartAppServer(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_RetrieveEnvironmentInfo() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.RetrieveEnvironmentInfoInput{
		InfoType:        aws.String("EnvironmentInfoType"), // Required
		EnvironmentID:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
	}
	resp, err := svc.RetrieveEnvironmentInfo(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_SwapEnvironmentCNAMEs() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.SwapEnvironmentCNAMEsInput{
		DestinationEnvironmentID:   aws.String("EnvironmentId"),
		DestinationEnvironmentName: aws.String("EnvironmentName"),
		SourceEnvironmentID:        aws.String("EnvironmentId"),
		SourceEnvironmentName:      aws.String("EnvironmentName"),
	}
	resp, err := svc.SwapEnvironmentCNAMEs(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_TerminateEnvironment() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.TerminateEnvironmentInput{
		EnvironmentID:      aws.String("EnvironmentId"),
		EnvironmentName:    aws.String("EnvironmentName"),
		TerminateResources: aws.Boolean(true),
	}
	resp, err := svc.TerminateEnvironment(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_UpdateApplication() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.UpdateApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		Description:     aws.String("Description"),
	}
	resp, err := svc.UpdateApplication(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_UpdateApplicationVersion() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.UpdateApplicationVersionInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		VersionLabel:    aws.String("VersionLabel"),    // Required
		Description:     aws.String("Description"),
	}
	resp, err := svc.UpdateApplicationVersion(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_UpdateConfigurationTemplate() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.UpdateConfigurationTemplateInput{
		ApplicationName: aws.String("ApplicationName"),           // Required
		TemplateName:    aws.String("ConfigurationTemplateName"), // Required
		Description:     aws.String("Description"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			&elasticbeanstalk.ConfigurationOptionSetting{ // Required
				Namespace:  aws.String("OptionNamespace"),
				OptionName: aws.String("ConfigurationOptionName"),
				Value:      aws.String("ConfigurationOptionValue"),
			},
			// More values...
		},
		OptionsToRemove: []*elasticbeanstalk.OptionSpecification{
			&elasticbeanstalk.OptionSpecification{ // Required
				Namespace:  aws.String("OptionNamespace"),
				OptionName: aws.String("ConfigurationOptionName"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateConfigurationTemplate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_UpdateEnvironment() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.UpdateEnvironmentInput{
		Description:     aws.String("Description"),
		EnvironmentID:   aws.String("EnvironmentId"),
		EnvironmentName: aws.String("EnvironmentName"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			&elasticbeanstalk.ConfigurationOptionSetting{ // Required
				Namespace:  aws.String("OptionNamespace"),
				OptionName: aws.String("ConfigurationOptionName"),
				Value:      aws.String("ConfigurationOptionValue"),
			},
			// More values...
		},
		OptionsToRemove: []*elasticbeanstalk.OptionSpecification{
			&elasticbeanstalk.OptionSpecification{ // Required
				Namespace:  aws.String("OptionNamespace"),
				OptionName: aws.String("ConfigurationOptionName"),
			},
			// More values...
		},
		TemplateName: aws.String("ConfigurationTemplateName"),
		Tier: &elasticbeanstalk.EnvironmentTier{
			Name:    aws.String("String"),
			Type:    aws.String("String"),
			Version: aws.String("String"),
		},
		VersionLabel: aws.String("VersionLabel"),
	}
	resp, err := svc.UpdateEnvironment(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticBeanstalk_ValidateConfigurationSettings() {
	svc := elasticbeanstalk.New(nil)

	params := &elasticbeanstalk.ValidateConfigurationSettingsInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{ // Required
			&elasticbeanstalk.ConfigurationOptionSetting{ // Required
				Namespace:  aws.String("OptionNamespace"),
				OptionName: aws.String("ConfigurationOptionName"),
				Value:      aws.String("ConfigurationOptionValue"),
			},
			// More values...
		},
		EnvironmentName: aws.String("EnvironmentName"),
		TemplateName:    aws.String("ConfigurationTemplateName"),
	}
	resp, err := svc.ValidateConfigurationSettings(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}