package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a new Step Functions Activity.
//
// Example:
//   activity := sfn.NewActivity(this, jsii.String("Activity"))
//
//   role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   })
//
//   activity.Grant(role, jsii.String("states:SendTaskSuccess"))
//
type Activity interface {
	awscdk.Resource
	IActivity
	// The ARN of the activity.
	ActivityArn() *string
	// The name of the activity.
	ActivityName() *string
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
	// Grant the given identity permissions on this Activity.
	Grant(identity awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Return the given named metric for this Activity.
	// Default: sum over 5 minutes.
	//
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity fails.
	// Default: sum over 5 minutes.
	//
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times the heartbeat times out for this activity.
	// Default: sum over 5 minutes.
	//
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the activity starts and the time it closes.
	// Default: average over 5 minutes.
	//
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is scheduled.
	// Default: sum over 5 minutes.
	//
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, for which the activity stays in the schedule state.
	// Default: average over 5 minutes.
	//
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is started.
	// Default: sum over 5 minutes.
	//
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity succeeds.
	// Default: sum over 5 minutes.
	//
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
	// Default: average over 5 minutes.
	//
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity times out.
	// Default: sum over 5 minutes.
	//
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Activity
type jsiiProxy_Activity struct {
	internal.Type__awscdkResource
	jsiiProxy_IActivity
}

func (j *jsiiProxy_Activity) ActivityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Activity) ActivityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Activity) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Activity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Activity) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Activity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewActivity(scope constructs.Construct, id *string, props *ActivityProps) Activity {
	_init_.Initialize()

	if err := validateNewActivityParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Activity{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Activity",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewActivity_Override(a Activity, scope constructs.Construct, id *string, props *ActivityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Activity",
		[]interface{}{scope, id, props},
		a,
	)
}

// Construct an Activity from an existing Activity ARN.
func Activity_FromActivityArn(scope constructs.Construct, id *string, activityArn *string) IActivity {
	_init_.Initialize()

	if err := validateActivity_FromActivityArnParameters(scope, id, activityArn); err != nil {
		panic(err)
	}
	var returns IActivity

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Activity",
		"fromActivityArn",
		[]interface{}{scope, id, activityArn},
		&returns,
	)

	return returns
}

// Construct an Activity from an existing Activity Name.
func Activity_FromActivityName(scope constructs.Construct, id *string, activityName *string) IActivity {
	_init_.Initialize()

	if err := validateActivity_FromActivityNameParameters(scope, id, activityName); err != nil {
		panic(err)
	}
	var returns IActivity

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Activity",
		"fromActivityName",
		[]interface{}{scope, id, activityName},
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
func Activity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActivity_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Activity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Activity_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateActivity_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Activity",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Activity_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateActivity_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Activity",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_Activity) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := a.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) GetResourceNameAttribute(nameAttr *string) *string {
	if err := a.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) Grant(identity awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := a.validateGrantParameters(identity); err != nil {
		panic(err)
	}
	args := []interface{}{identity}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

