// +build !ignore_autogenerated

/*
Copyright 2021 - now, the original author or authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package prometheus

import (
	"github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConfig) DeepCopyInto(out *MonitoringConfig) {
	*out = *in
	if in.TlsConfig != nil {
		in, out := &in.TlsConfig, &out.TlsConfig
		*out = new(v1.TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(v1.BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	in.BearerTokenSecret.DeepCopyInto(&out.BearerTokenSecret)
	if in.ReLabelings != nil {
		in, out := &in.ReLabelings, &out.ReLabelings
		*out = make([]*v1.RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.RelabelConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.MetricReLabelings != nil {
		in, out := &in.MetricReLabelings, &out.MetricReLabelings
		*out = make([]*v1.RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.RelabelConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConfig.
func (in *MonitoringConfig) DeepCopy() *MonitoringConfig {
	if in == nil {
		return nil
	}
	out := new(MonitoringConfig)
	in.DeepCopyInto(out)
	return out
}
