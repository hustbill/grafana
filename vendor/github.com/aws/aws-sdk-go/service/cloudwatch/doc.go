// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudwatch provides the client and types for making API
// requests to Amazon CloudWatch.
//
// Amazon CloudWatch monitors your Amazon Web Services (AWS) resources and the
// applications you run on AWS in real time. You can use CloudWatch to collect
// and track metrics, which are the variables you want to measure for your resources
// and applications.
//
// CloudWatch alarms send notifications or automatically change the resources
// you are monitoring based on rules that you define. For example, you can monitor
// the CPU usage and disk reads and writes of your Amazon EC2 instances. Then,
// use this data to determine whether you should launch additional instances
// to handle increased load. You can also use this data to stop under-used instances
// to save money.
//
// In addition to monitoring the built-in metrics that come with AWS, you can
// monitor your own custom metrics. With CloudWatch, you gain system-wide visibility
// into resource utilization, application performance, and operational health.
//
// See https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01 for more information on this service.
//
// See cloudwatch package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatch/
//
// Using the Client
//
// To use the client for Amazon CloudWatch you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := cloudwatch.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon CloudWatch client CloudWatch for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatch/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.DeleteAlarms(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("DeleteAlarms result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.DeleteAlarmsWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package cloudwatch
