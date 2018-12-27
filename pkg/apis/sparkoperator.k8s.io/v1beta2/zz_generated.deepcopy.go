// +build !ignore_autogenerated

// Code generated by k8s code-generator DO NOT EDIT.

/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta2

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationState) DeepCopyInto(out *ApplicationState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationState.
func (in *ApplicationState) DeepCopy() *ApplicationState {
	if in == nil {
		return nil
	}
	out := new(ApplicationState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchSchedulerConfiguration) DeepCopyInto(out *BatchSchedulerConfiguration) {
	*out = *in
	if in.Queue != nil {
		in, out := &in.Queue, &out.Queue
		*out = new(string)
		**out = **in
	}
	if in.PriorityClassName != nil {
		in, out := &in.PriorityClassName, &out.PriorityClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchSchedulerConfiguration.
func (in *BatchSchedulerConfiguration) DeepCopy() *BatchSchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(BatchSchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dependencies) DeepCopyInto(out *Dependencies) {
	*out = *in
	if in.Jars != nil {
		in, out := &in.Jars, &out.Jars
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PyFiles != nil {
		in, out := &in.PyFiles, &out.PyFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dependencies.
func (in *Dependencies) DeepCopy() *Dependencies {
	if in == nil {
		return nil
	}
	out := new(Dependencies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverInfo) DeepCopyInto(out *DriverInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverInfo.
func (in *DriverInfo) DeepCopy() *DriverInfo {
	if in == nil {
		return nil
	}
	out := new(DriverInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverSpec) DeepCopyInto(out *DriverSpec) {
	*out = *in
	in.SparkPodSpec.DeepCopyInto(&out.SparkPodSpec)
	if in.PodName != nil {
		in, out := &in.PodName, &out.PodName
		*out = new(string)
		**out = **in
	}
	if in.CoreRequest != nil {
		in, out := &in.CoreRequest, &out.CoreRequest
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.JavaOptions != nil {
		in, out := &in.JavaOptions, &out.JavaOptions
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverSpec.
func (in *DriverSpec) DeepCopy() *DriverSpec {
	if in == nil {
		return nil
	}
	out := new(DriverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutorSpec) DeepCopyInto(out *ExecutorSpec) {
	*out = *in
	in.SparkPodSpec.DeepCopyInto(&out.SparkPodSpec)
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(int32)
		**out = **in
	}
	if in.CoreRequest != nil {
		in, out := &in.CoreRequest, &out.CoreRequest
		*out = new(string)
		**out = **in
	}
	if in.JavaOptions != nil {
		in, out := &in.JavaOptions, &out.JavaOptions
		*out = new(string)
		**out = **in
	}
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutorSpec.
func (in *ExecutorSpec) DeepCopy() *ExecutorSpec {
	if in == nil {
		return nil
	}
	out := new(ExecutorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPUSpec) DeepCopyInto(out *GPUSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPUSpec.
func (in *GPUSpec) DeepCopy() *GPUSpec {
	if in == nil {
		return nil
	}
	out := new(GPUSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringSpec) DeepCopyInto(out *MonitoringSpec) {
	*out = *in
	if in.MetricsProperties != nil {
		in, out := &in.MetricsProperties, &out.MetricsProperties
		*out = new(string)
		**out = **in
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringSpec.
func (in *MonitoringSpec) DeepCopy() *MonitoringSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameKey) DeepCopyInto(out *NameKey) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameKey.
func (in *NameKey) DeepCopy() *NameKey {
	if in == nil {
		return nil
	}
	out := new(NameKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamePath) DeepCopyInto(out *NamePath) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamePath.
func (in *NamePath) DeepCopy() *NamePath {
	if in == nil {
		return nil
	}
	out := new(NamePath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusSpec) DeepCopyInto(out *PrometheusSpec) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.ConfigFile != nil {
		in, out := &in.ConfigFile, &out.ConfigFile
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusSpec.
func (in *PrometheusSpec) DeepCopy() *PrometheusSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestartPolicy) DeepCopyInto(out *RestartPolicy) {
	*out = *in
	if in.OnSubmissionFailureRetries != nil {
		in, out := &in.OnSubmissionFailureRetries, &out.OnSubmissionFailureRetries
		*out = new(int32)
		**out = **in
	}
	if in.OnFailureRetries != nil {
		in, out := &in.OnFailureRetries, &out.OnFailureRetries
		*out = new(int32)
		**out = **in
	}
	if in.OnSubmissionFailureRetryInterval != nil {
		in, out := &in.OnSubmissionFailureRetryInterval, &out.OnSubmissionFailureRetryInterval
		*out = new(int64)
		**out = **in
	}
	if in.OnFailureRetryInterval != nil {
		in, out := &in.OnFailureRetryInterval, &out.OnFailureRetryInterval
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestartPolicy.
func (in *RestartPolicy) DeepCopy() *RestartPolicy {
	if in == nil {
		return nil
	}
	out := new(RestartPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledSparkApplication) DeepCopyInto(out *ScheduledSparkApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledSparkApplication.
func (in *ScheduledSparkApplication) DeepCopy() *ScheduledSparkApplication {
	if in == nil {
		return nil
	}
	out := new(ScheduledSparkApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledSparkApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledSparkApplicationList) DeepCopyInto(out *ScheduledSparkApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScheduledSparkApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledSparkApplicationList.
func (in *ScheduledSparkApplicationList) DeepCopy() *ScheduledSparkApplicationList {
	if in == nil {
		return nil
	}
	out := new(ScheduledSparkApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledSparkApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledSparkApplicationSpec) DeepCopyInto(out *ScheduledSparkApplicationSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
	if in.SuccessfulRunHistoryLimit != nil {
		in, out := &in.SuccessfulRunHistoryLimit, &out.SuccessfulRunHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.FailedRunHistoryLimit != nil {
		in, out := &in.FailedRunHistoryLimit, &out.FailedRunHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledSparkApplicationSpec.
func (in *ScheduledSparkApplicationSpec) DeepCopy() *ScheduledSparkApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduledSparkApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledSparkApplicationStatus) DeepCopyInto(out *ScheduledSparkApplicationStatus) {
	*out = *in
	in.LastRun.DeepCopyInto(&out.LastRun)
	in.NextRun.DeepCopyInto(&out.NextRun)
	if in.PastSuccessfulRunNames != nil {
		in, out := &in.PastSuccessfulRunNames, &out.PastSuccessfulRunNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PastFailedRunNames != nil {
		in, out := &in.PastFailedRunNames, &out.PastFailedRunNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledSparkApplicationStatus.
func (in *ScheduledSparkApplicationStatus) DeepCopy() *ScheduledSparkApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduledSparkApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretInfo) DeepCopyInto(out *SecretInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretInfo.
func (in *SecretInfo) DeepCopy() *SecretInfo {
	if in == nil {
		return nil
	}
	out := new(SecretInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkApplication) DeepCopyInto(out *SparkApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkApplication.
func (in *SparkApplication) DeepCopy() *SparkApplication {
	if in == nil {
		return nil
	}
	out := new(SparkApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SparkApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkApplicationList) DeepCopyInto(out *SparkApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SparkApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkApplicationList.
func (in *SparkApplicationList) DeepCopy() *SparkApplicationList {
	if in == nil {
		return nil
	}
	out := new(SparkApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SparkApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkApplicationSpec) DeepCopyInto(out *SparkApplicationSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(string)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MainClass != nil {
		in, out := &in.MainClass, &out.MainClass
		*out = new(string)
		**out = **in
	}
	if in.MainApplicationFile != nil {
		in, out := &in.MainApplicationFile, &out.MainApplicationFile
		*out = new(string)
		**out = **in
	}
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SparkConf != nil {
		in, out := &in.SparkConf, &out.SparkConf
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HadoopConf != nil {
		in, out := &in.HadoopConf, &out.HadoopConf
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SparkConfigMap != nil {
		in, out := &in.SparkConfigMap, &out.SparkConfigMap
		*out = new(string)
		**out = **in
	}
	if in.HadoopConfigMap != nil {
		in, out := &in.HadoopConfigMap, &out.HadoopConfigMap
		*out = new(string)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Driver.DeepCopyInto(&out.Driver)
	in.Executor.DeepCopyInto(&out.Executor)
	in.Deps.DeepCopyInto(&out.Deps)
	in.RestartPolicy.DeepCopyInto(&out.RestartPolicy)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FailureRetries != nil {
		in, out := &in.FailureRetries, &out.FailureRetries
		*out = new(int32)
		**out = **in
	}
	if in.RetryInterval != nil {
		in, out := &in.RetryInterval, &out.RetryInterval
		*out = new(int64)
		**out = **in
	}
	if in.PythonVersion != nil {
		in, out := &in.PythonVersion, &out.PythonVersion
		*out = new(string)
		**out = **in
	}
	if in.MemoryOverheadFactor != nil {
		in, out := &in.MemoryOverheadFactor, &out.MemoryOverheadFactor
		*out = new(string)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(MonitoringSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.BatchScheduler != nil {
		in, out := &in.BatchScheduler, &out.BatchScheduler
		*out = new(string)
		**out = **in
	}
	if in.BatchSchedulerOptions != nil {
		in, out := &in.BatchSchedulerOptions, &out.BatchSchedulerOptions
		*out = new(BatchSchedulerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.TimeToLiveSeconds != nil {
		in, out := &in.TimeToLiveSeconds, &out.TimeToLiveSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkApplicationSpec.
func (in *SparkApplicationSpec) DeepCopy() *SparkApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(SparkApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkApplicationStatus) DeepCopyInto(out *SparkApplicationStatus) {
	*out = *in
	in.SubmissionTime.DeepCopyInto(&out.SubmissionTime)
	in.TerminationTime.DeepCopyInto(&out.TerminationTime)
	out.DriverInfo = in.DriverInfo
	out.AppState = in.AppState
	if in.ExecutorState != nil {
		in, out := &in.ExecutorState, &out.ExecutorState
		*out = make(map[string]ExecutorState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkApplicationStatus.
func (in *SparkApplicationStatus) DeepCopy() *SparkApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(SparkApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkPodSpec) DeepCopyInto(out *SparkPodSpec) {
	*out = *in
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(int32)
		**out = **in
	}
	if in.CoreLimit != nil {
		in, out := &in.CoreLimit, &out.CoreLimit
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
	if in.MemoryOverhead != nil {
		in, out := &in.MemoryOverhead, &out.MemoryOverhead
		*out = new(string)
		**out = **in
	}
	if in.GPU != nil {
		in, out := &in.GPU, &out.GPU
		*out = new(GPUSpec)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]NamePath, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]SecretInfo, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnvSecretKeyRefs != nil {
		in, out := &in.EnvSecretKeyRefs, &out.EnvSecretKeyRefs
		*out = make(map[string]NameKey, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContenxt != nil {
		in, out := &in.SecurityContenxt, &out.SecurityContenxt
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.SchedulerName != nil {
		in, out := &in.SchedulerName, &out.SchedulerName
		*out = new(string)
		**out = **in
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DNSConfig != nil {
		in, out := &in.DNSConfig, &out.DNSConfig
		*out = new(v1.PodDNSConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkPodSpec.
func (in *SparkPodSpec) DeepCopy() *SparkPodSpec {
	if in == nil {
		return nil
	}
	out := new(SparkPodSpec)
	in.DeepCopyInto(out)
	return out
}
