package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a REST API in Amazon API Gateway, created with an OpenAPI specification.
//
// Some properties normally accessible on.
//
// Example:
//   var integration integration
//
//
//   api := apigateway.NewSpecRestApi(this, jsii.String("books-api"), &SpecRestApiProps{
//   	ApiDefinition: apigateway.ApiDefinition_FromAsset(jsii.String("path-to-file.json")),
//   })
//
//   booksResource := api.Root.AddResource(jsii.String("books"))
//   booksResource.AddMethod(jsii.String("GET"), integration)
//
// See:  `RestApi` - such as the description -
// must be declared in the specification. All Resources and Methods need to be defined as
// part of the OpenAPI specification file, and cannot be added via the CDK.
//
// By default, the API will automatically be deployed and accessible from a
// public endpoint.
//
type SpecRestApi interface {
	RestApiBase
	CloudWatchAccount() CfnAccount
	SetCloudWatchAccount(val CfnAccount)
	// API Gateway stage that points to the latest deployment (if defined).
	//
	// If `deploy` is disabled, you will need to explicitly assign this value in order to
	// set up integrations.
	DeploymentStage() Stage
	SetDeploymentStage(val Stage)
	// The first domain name mapped to this API, if defined through the `domainName` configuration prop, or added via `addDomainName`.
	DomainName() DomainName
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// API Gateway deployment that represents the latest changes of the API.
	//
	// This resource will be automatically updated every time the REST API model changes.
	// This will be undefined if `deploy` is false.
	LatestDeployment() Deployment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The ID of this API Gateway RestApi.
	RestApiId() *string
	// A human friendly name for this Rest API.
	//
	// Note that this is different from `restApiId`.
	RestApiName() *string
	// The resource ID of the root resource.
	RestApiRootResourceId() *string
	// Represents the root resource of this API endpoint ('/').
	//
	// Resources and Methods are added to this resource.
	Root() IResource
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The deployed root URL of this REST API.
	Url() *string
	// Add an ApiKey to the deploymentStage.
	AddApiKey(id *string, options *ApiKeyOptions) IApiKey
	// Defines an API Gateway domain name and maps it to this API.
	AddDomainName(id *string, options *DomainNameOptions) DomainName
	// Adds a new gateway response.
	AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse
	// Adds a usage plan.
	AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Gets the "execute-api" ARN.
	ArnForExecuteApi(method *string, path *string, stage *string) *string
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns the given named metric for this API.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the API cache in a given period.
	//
	// Default: sum over 5 minutes.
	MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the backend in a given period, when API caching is enabled.
	//
	// Default: sum over 5 minutes.
	MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of client-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	//
	// Default: sample count over 5 minutes.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	//
	// Default: average over 5 minutes.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	//
	// Default: average over 5 minutes.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	// Returns the URL for an HTTP path.
	//
	// Fails if `deploymentStage` is not set either by `deploy` or explicitly.
	UrlForPath(path *string) *string
}

// The jsii proxy struct for SpecRestApi
type jsiiProxy_SpecRestApi struct {
	jsiiProxy_RestApiBase
}

func (j *jsiiProxy_SpecRestApi) CloudWatchAccount() CfnAccount {
	var returns CfnAccount
	_jsii_.Get(
		j,
		"cloudWatchAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) DeploymentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"deploymentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) DomainName() DomainName {
	var returns DomainName
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) LatestDeployment() Deployment {
	var returns Deployment
	_jsii_.Get(
		j,
		"latestDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) RestApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) RestApiRootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiRootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) Root() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpecRestApi) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewSpecRestApi(scope constructs.Construct, id *string, props *SpecRestApiProps) SpecRestApi {
	_init_.Initialize()

	if err := validateNewSpecRestApiParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpecRestApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.SpecRestApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSpecRestApi_Override(s SpecRestApi, scope constructs.Construct, id *string, props *SpecRestApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.SpecRestApi",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SpecRestApi)SetCloudWatchAccount(val CfnAccount) {
	_jsii_.Set(
		j,
		"cloudWatchAccount",
		val,
	)
}

func (j *jsiiProxy_SpecRestApi)SetDeploymentStage(val Stage) {
	if err := j.validateSetDeploymentStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentStage",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SpecRestApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpecRestApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.SpecRestApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func SpecRestApi_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSpecRestApi_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.SpecRestApi",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func SpecRestApi_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSpecRestApi_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.SpecRestApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) AddApiKey(id *string, options *ApiKeyOptions) IApiKey {
	if err := s.validateAddApiKeyParameters(id, options); err != nil {
		panic(err)
	}
	var returns IApiKey

	_jsii_.Invoke(
		s,
		"addApiKey",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) AddDomainName(id *string, options *DomainNameOptions) DomainName {
	if err := s.validateAddDomainNameParameters(id, options); err != nil {
		panic(err)
	}
	var returns DomainName

	_jsii_.Invoke(
		s,
		"addDomainName",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse {
	if err := s.validateAddGatewayResponseParameters(id, options); err != nil {
		panic(err)
	}
	var returns GatewayResponse

	_jsii_.Invoke(
		s,
		"addGatewayResponse",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan {
	if err := s.validateAddUsagePlanParameters(id, props); err != nil {
		panic(err)
	}
	var returns UsagePlan

	_jsii_.Invoke(
		s,
		"addUsagePlan",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SpecRestApi) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricCacheHitCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCacheHitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricCacheMissCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCacheMissCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricClientErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricIntegrationLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricServerErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpecRestApi) UrlForPath(path *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"urlForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

