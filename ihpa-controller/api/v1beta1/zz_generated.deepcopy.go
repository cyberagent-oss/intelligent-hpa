// +build !ignore_autogenerated

/*
Copyright 2020 SIA Platform Team.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChangePointDetectionConfig) DeepCopyInto(out *ChangePointDetectionConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChangePointDetectionConfig.
func (in *ChangePointDetectionConfig) DeepCopy() *ChangePointDetectionConfig {
	if in == nil {
		return nil
	}
	out := new(ChangePointDetectionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogProviderSource) DeepCopyInto(out *DatadogProviderSource) {
	*out = *in
	if in.KeysFrom != nil {
		in, out := &in.KeysFrom, &out.KeysFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogProviderSource.
func (in *DatadogProviderSource) DeepCopy() *DatadogProviderSource {
	if in == nil {
		return nil
	}
	out := new(DatadogProviderSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FittingJob) DeepCopyInto(out *FittingJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FittingJob.
func (in *FittingJob) DeepCopy() *FittingJob {
	if in == nil {
		return nil
	}
	out := new(FittingJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FittingJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FittingJobConfig) DeepCopyInto(out *FittingJobConfig) {
	*out = *in
	out.ChangePointDetectionConfig = in.ChangePointDetectionConfig
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FittingJobConfig.
func (in *FittingJobConfig) DeepCopy() *FittingJobConfig {
	if in == nil {
		return nil
	}
	out := new(FittingJobConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FittingJobList) DeepCopyInto(out *FittingJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FittingJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FittingJobList.
func (in *FittingJobList) DeepCopy() *FittingJobList {
	if in == nil {
		return nil
	}
	out := new(FittingJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FittingJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FittingJobSpec) DeepCopyInto(out *FittingJobSpec) {
	*out = *in
	out.ChangePointDetectionConfig = in.ChangePointDetectionConfig
	out.DataConfigMap = in.DataConfigMap
	in.JobTemplate.DeepCopyInto(&out.JobTemplate)
	in.TargetMetric.DeepCopyInto(&out.TargetMetric)
	in.Provider.DeepCopyInto(&out.Provider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FittingJobSpec.
func (in *FittingJobSpec) DeepCopy() *FittingJobSpec {
	if in == nil {
		return nil
	}
	out := new(FittingJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FittingJobStatus) DeepCopyInto(out *FittingJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FittingJobStatus.
func (in *FittingJobStatus) DeepCopy() *FittingJobStatus {
	if in == nil {
		return nil
	}
	out := new(FittingJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerTemplateSpec) DeepCopyInto(out *HorizontalPodAutoscalerTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerTemplateSpec.
func (in *HorizontalPodAutoscalerTemplateSpec) DeepCopy() *HorizontalPodAutoscalerTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntelligentHorizontalPodAutoscaler) DeepCopyInto(out *IntelligentHorizontalPodAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntelligentHorizontalPodAutoscaler.
func (in *IntelligentHorizontalPodAutoscaler) DeepCopy() *IntelligentHorizontalPodAutoscaler {
	if in == nil {
		return nil
	}
	out := new(IntelligentHorizontalPodAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IntelligentHorizontalPodAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntelligentHorizontalPodAutoscalerList) DeepCopyInto(out *IntelligentHorizontalPodAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IntelligentHorizontalPodAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntelligentHorizontalPodAutoscalerList.
func (in *IntelligentHorizontalPodAutoscalerList) DeepCopy() *IntelligentHorizontalPodAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(IntelligentHorizontalPodAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IntelligentHorizontalPodAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntelligentHorizontalPodAutoscalerSpec) DeepCopyInto(out *IntelligentHorizontalPodAutoscalerSpec) {
	*out = *in
	in.HorizontalPodAutoscalerTemplate.DeepCopyInto(&out.HorizontalPodAutoscalerTemplate)
	in.FittingJobConfig.DeepCopyInto(&out.FittingJobConfig)
	in.MetricProvider.DeepCopyInto(&out.MetricProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntelligentHorizontalPodAutoscalerSpec.
func (in *IntelligentHorizontalPodAutoscalerSpec) DeepCopy() *IntelligentHorizontalPodAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(IntelligentHorizontalPodAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntelligentHorizontalPodAutoscalerStatus) DeepCopyInto(out *IntelligentHorizontalPodAutoscalerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntelligentHorizontalPodAutoscalerStatus.
func (in *IntelligentHorizontalPodAutoscalerStatus) DeepCopy() *IntelligentHorizontalPodAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(IntelligentHorizontalPodAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricProvider) DeepCopyInto(out *MetricProvider) {
	*out = *in
	in.ProviderSource.DeepCopyInto(&out.ProviderSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricProvider.
func (in *MetricProvider) DeepCopy() *MetricProvider {
	if in == nil {
		return nil
	}
	out := new(MetricProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusProviderSource) DeepCopyInto(out *PrometheusProviderSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusProviderSource.
func (in *PrometheusProviderSource) DeepCopy() *PrometheusProviderSource {
	if in == nil {
		return nil
	}
	out := new(PrometheusProviderSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSource) DeepCopyInto(out *ProviderSource) {
	*out = *in
	if in.Datadog != nil {
		in, out := &in.Datadog, &out.Datadog
		*out = new(DatadogProviderSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusProviderSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSource.
func (in *ProviderSource) DeepCopy() *ProviderSource {
	if in == nil {
		return nil
	}
	out := new(ProviderSource)
	in.DeepCopyInto(out)
	return out
}
