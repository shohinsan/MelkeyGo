package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a user pool domain.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("UserPool"), &UserPoolProps{
//   })
//   client := userpool.addClient(jsii.String("Client"), &UserPoolClientOptions{
//   	// ...
//   	OAuth: &OAuthSettings{
//   		Flows: &OAuthFlows{
//   			ImplicitCodeGrant: jsii.Boolean(true),
//   		},
//   		CallbackUrls: []*string{
//   			jsii.String("https://myapp.com/home"),
//   			jsii.String("https://myapp.com/users"),
//   		},
//   	},
//   })
//   domain := userpool.addDomain(jsii.String("Domain"), &UserPoolDomainOptions{
//   })
//   signInUrl := domain.SignInUrl(client, &SignInUrlOptions{
//   	RedirectUri: jsii.String("https://myapp.com/home"),
//   })
//
type UserPoolDomain interface {
	awscdk.Resource
	IUserPoolDomain
	// The domain name of the CloudFront distribution associated with the user pool domain.
	CloudFrontDomainName() *string
	// The domain that was specified to be created.
	//
	// If `customDomain` was selected, this holds the full domain name that was specified.
	// If the `cognitoDomain` was used, it contains the prefix to the Cognito hosted domain.
	DomainName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
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
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
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
	// The URL to the hosted UI associated with this domain.
	BaseUrl(options *BaseUrlOptions) *string
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
	// The URL to the sign in page in this domain using a specific UserPoolClient.
	SignInUrl(client UserPoolClient, options *SignInUrlOptions) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for UserPoolDomain
type jsiiProxy_UserPoolDomain struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPoolDomain
}

func (j *jsiiProxy_UserPoolDomain) CloudFrontDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFrontDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolDomain) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolDomain) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolDomain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewUserPoolDomain(scope constructs.Construct, id *string, props *UserPoolDomainProps) UserPoolDomain {
	_init_.Initialize()

	if err := validateNewUserPoolDomainParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserPoolDomain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.UserPoolDomain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewUserPoolDomain_Override(u UserPoolDomain, scope constructs.Construct, id *string, props *UserPoolDomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.UserPoolDomain",
		[]interface{}{scope, id, props},
		u,
	)
}

// Import a UserPoolDomain given its domain name.
func UserPoolDomain_FromDomainName(scope constructs.Construct, id *string, userPoolDomainName *string) IUserPoolDomain {
	_init_.Initialize()

	if err := validateUserPoolDomain_FromDomainNameParameters(scope, id, userPoolDomainName); err != nil {
		panic(err)
	}
	var returns IUserPoolDomain

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolDomain",
		"fromDomainName",
		[]interface{}{scope, id, userPoolDomainName},
		&returns,
	)

	return returns
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
func UserPoolDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUserPoolDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func UserPoolDomain_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateUserPoolDomain_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolDomain",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func UserPoolDomain_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateUserPoolDomain_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolDomain",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := u.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPoolDomain) BaseUrl(options *BaseUrlOptions) *string {
	if err := u.validateBaseUrlParameters(options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"baseUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := u.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) GetResourceNameAttribute(nameAttr *string) *string {
	if err := u.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) SignInUrl(client UserPoolClient, options *SignInUrlOptions) *string {
	if err := u.validateSignInUrlParameters(client, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"signInUrl",
		[]interface{}{client, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
