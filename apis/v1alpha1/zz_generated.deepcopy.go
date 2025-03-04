//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alarm) DeepCopyInto(out *Alarm) {
	*out = *in
	if in.AlarmARN != nil {
		in, out := &in.AlarmARN, &out.AlarmARN
		*out = new(string)
		**out = **in
	}
	if in.AlarmName != nil {
		in, out := &in.AlarmName, &out.AlarmName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alarm.
func (in *Alarm) DeepCopy() *Alarm {
	if in == nil {
		return nil
	}
	out := new(Alarm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedMetricSpecification) DeepCopyInto(out *CustomizedMetricSpecification) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]*MetricDimension, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MetricDimension)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedMetricSpecification.
func (in *CustomizedMetricSpecification) DeepCopy() *CustomizedMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(CustomizedMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricDimension) DeepCopyInto(out *MetricDimension) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricDimension.
func (in *MetricDimension) DeepCopy() *MetricDimension {
	if in == nil {
		return nil
	}
	out := new(MetricDimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredefinedMetricSpecification) DeepCopyInto(out *PredefinedMetricSpecification) {
	*out = *in
	if in.PredefinedMetricType != nil {
		in, out := &in.PredefinedMetricType, &out.PredefinedMetricType
		*out = new(string)
		**out = **in
	}
	if in.ResourceLabel != nil {
		in, out := &in.ResourceLabel, &out.ResourceLabel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredefinedMetricSpecification.
func (in *PredefinedMetricSpecification) DeepCopy() *PredefinedMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(PredefinedMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTarget) DeepCopyInto(out *ScalableTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTarget.
func (in *ScalableTarget) DeepCopy() *ScalableTarget {
	if in == nil {
		return nil
	}
	out := new(ScalableTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalableTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetAction) DeepCopyInto(out *ScalableTargetAction) {
	*out = *in
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(int64)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetAction.
func (in *ScalableTargetAction) DeepCopy() *ScalableTargetAction {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetList) DeepCopyInto(out *ScalableTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScalableTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetList.
func (in *ScalableTargetList) DeepCopy() *ScalableTargetList {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalableTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetSpec) DeepCopyInto(out *ScalableTargetSpec) {
	*out = *in
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(int64)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(int64)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.RoleARN != nil {
		in, out := &in.RoleARN, &out.RoleARN
		*out = new(string)
		**out = **in
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
	if in.SuspendedState != nil {
		in, out := &in.SuspendedState, &out.SuspendedState
		*out = new(SuspendedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetSpec.
func (in *ScalableTargetSpec) DeepCopy() *ScalableTargetSpec {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetStatus) DeepCopyInto(out *ScalableTargetStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.LastModifiedTime != nil {
		in, out := &in.LastModifiedTime, &out.LastModifiedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetStatus.
func (in *ScalableTargetStatus) DeepCopy() *ScalableTargetStatus {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTarget_SDK) DeepCopyInto(out *ScalableTarget_SDK) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(int64)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(int64)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.RoleARN != nil {
		in, out := &in.RoleARN, &out.RoleARN
		*out = new(string)
		**out = **in
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
	if in.SuspendedState != nil {
		in, out := &in.SuspendedState, &out.SuspendedState
		*out = new(SuspendedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTarget_SDK.
func (in *ScalableTarget_SDK) DeepCopy() *ScalableTarget_SDK {
	if in == nil {
		return nil
	}
	out := new(ScalableTarget_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingActivity) DeepCopyInto(out *ScalingActivity) {
	*out = *in
	if in.ActivityID != nil {
		in, out := &in.ActivityID, &out.ActivityID
		*out = new(string)
		**out = **in
	}
	if in.Cause != nil {
		in, out := &in.Cause, &out.Cause
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingActivity.
func (in *ScalingActivity) DeepCopy() *ScalingActivity {
	if in == nil {
		return nil
	}
	out := new(ScalingActivity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicy) DeepCopyInto(out *ScalingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicy.
func (in *ScalingPolicy) DeepCopy() *ScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyList) DeepCopyInto(out *ScalingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScalingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyList.
func (in *ScalingPolicyList) DeepCopy() *ScalingPolicyList {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicySpec) DeepCopyInto(out *ScalingPolicySpec) {
	*out = *in
	if in.PolicyName != nil {
		in, out := &in.PolicyName, &out.PolicyName
		*out = new(string)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
	if in.StepScalingPolicyConfiguration != nil {
		in, out := &in.StepScalingPolicyConfiguration, &out.StepScalingPolicyConfiguration
		*out = new(StepScalingPolicyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetTrackingScalingPolicyConfiguration != nil {
		in, out := &in.TargetTrackingScalingPolicyConfiguration, &out.TargetTrackingScalingPolicyConfiguration
		*out = new(TargetTrackingScalingPolicyConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicySpec.
func (in *ScalingPolicySpec) DeepCopy() *ScalingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyStatus) DeepCopyInto(out *ScalingPolicyStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Alarms != nil {
		in, out := &in.Alarms, &out.Alarms
		*out = make([]*Alarm, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Alarm)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.LastModifiedTime != nil {
		in, out := &in.LastModifiedTime, &out.LastModifiedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyStatus.
func (in *ScalingPolicyStatus) DeepCopy() *ScalingPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicy_SDK) DeepCopyInto(out *ScalingPolicy_SDK) {
	*out = *in
	if in.Alarms != nil {
		in, out := &in.Alarms, &out.Alarms
		*out = make([]*Alarm, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Alarm)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.PolicyARN != nil {
		in, out := &in.PolicyARN, &out.PolicyARN
		*out = new(string)
		**out = **in
	}
	if in.PolicyName != nil {
		in, out := &in.PolicyName, &out.PolicyName
		*out = new(string)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
	if in.StepScalingPolicyConfiguration != nil {
		in, out := &in.StepScalingPolicyConfiguration, &out.StepScalingPolicyConfiguration
		*out = new(StepScalingPolicyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetTrackingScalingPolicyConfiguration != nil {
		in, out := &in.TargetTrackingScalingPolicyConfiguration, &out.TargetTrackingScalingPolicyConfiguration
		*out = new(TargetTrackingScalingPolicyConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicy_SDK.
func (in *ScalingPolicy_SDK) DeepCopy() *ScalingPolicy_SDK {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicy_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledAction) DeepCopyInto(out *ScheduledAction) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.ScheduledActionARN != nil {
		in, out := &in.ScheduledActionARN, &out.ScheduledActionARN
		*out = new(string)
		**out = **in
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledAction.
func (in *ScheduledAction) DeepCopy() *ScheduledAction {
	if in == nil {
		return nil
	}
	out := new(ScheduledAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepAdjustment) DeepCopyInto(out *StepAdjustment) {
	*out = *in
	if in.MetricIntervalLowerBound != nil {
		in, out := &in.MetricIntervalLowerBound, &out.MetricIntervalLowerBound
		*out = new(float64)
		**out = **in
	}
	if in.MetricIntervalUpperBound != nil {
		in, out := &in.MetricIntervalUpperBound, &out.MetricIntervalUpperBound
		*out = new(float64)
		**out = **in
	}
	if in.ScalingAdjustment != nil {
		in, out := &in.ScalingAdjustment, &out.ScalingAdjustment
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepAdjustment.
func (in *StepAdjustment) DeepCopy() *StepAdjustment {
	if in == nil {
		return nil
	}
	out := new(StepAdjustment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepScalingPolicyConfiguration) DeepCopyInto(out *StepScalingPolicyConfiguration) {
	*out = *in
	if in.AdjustmentType != nil {
		in, out := &in.AdjustmentType, &out.AdjustmentType
		*out = new(string)
		**out = **in
	}
	if in.Cooldown != nil {
		in, out := &in.Cooldown, &out.Cooldown
		*out = new(int64)
		**out = **in
	}
	if in.MetricAggregationType != nil {
		in, out := &in.MetricAggregationType, &out.MetricAggregationType
		*out = new(string)
		**out = **in
	}
	if in.MinAdjustmentMagnitude != nil {
		in, out := &in.MinAdjustmentMagnitude, &out.MinAdjustmentMagnitude
		*out = new(int64)
		**out = **in
	}
	if in.StepAdjustments != nil {
		in, out := &in.StepAdjustments, &out.StepAdjustments
		*out = make([]*StepAdjustment, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(StepAdjustment)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepScalingPolicyConfiguration.
func (in *StepScalingPolicyConfiguration) DeepCopy() *StepScalingPolicyConfiguration {
	if in == nil {
		return nil
	}
	out := new(StepScalingPolicyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuspendedState) DeepCopyInto(out *SuspendedState) {
	*out = *in
	if in.DynamicScalingInSuspended != nil {
		in, out := &in.DynamicScalingInSuspended, &out.DynamicScalingInSuspended
		*out = new(bool)
		**out = **in
	}
	if in.DynamicScalingOutSuspended != nil {
		in, out := &in.DynamicScalingOutSuspended, &out.DynamicScalingOutSuspended
		*out = new(bool)
		**out = **in
	}
	if in.ScheduledScalingSuspended != nil {
		in, out := &in.ScheduledScalingSuspended, &out.ScheduledScalingSuspended
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuspendedState.
func (in *SuspendedState) DeepCopy() *SuspendedState {
	if in == nil {
		return nil
	}
	out := new(SuspendedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetTrackingScalingPolicyConfiguration) DeepCopyInto(out *TargetTrackingScalingPolicyConfiguration) {
	*out = *in
	if in.CustomizedMetricSpecification != nil {
		in, out := &in.CustomizedMetricSpecification, &out.CustomizedMetricSpecification
		*out = new(CustomizedMetricSpecification)
		(*in).DeepCopyInto(*out)
	}
	if in.DisableScaleIn != nil {
		in, out := &in.DisableScaleIn, &out.DisableScaleIn
		*out = new(bool)
		**out = **in
	}
	if in.PredefinedMetricSpecification != nil {
		in, out := &in.PredefinedMetricSpecification, &out.PredefinedMetricSpecification
		*out = new(PredefinedMetricSpecification)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaleInCooldown != nil {
		in, out := &in.ScaleInCooldown, &out.ScaleInCooldown
		*out = new(int64)
		**out = **in
	}
	if in.ScaleOutCooldown != nil {
		in, out := &in.ScaleOutCooldown, &out.ScaleOutCooldown
		*out = new(int64)
		**out = **in
	}
	if in.TargetValue != nil {
		in, out := &in.TargetValue, &out.TargetValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetTrackingScalingPolicyConfiguration.
func (in *TargetTrackingScalingPolicyConfiguration) DeepCopy() *TargetTrackingScalingPolicyConfiguration {
	if in == nil {
		return nil
	}
	out := new(TargetTrackingScalingPolicyConfiguration)
	in.DeepCopyInto(out)
	return out
}
