//
//Copyright 2022 The KEDA Authors
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: metrics.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1beta1 "k8s.io/metrics/pkg/apis/external_metrics/v1beta1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScaledObjectRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace  string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	MetricName string `protobuf:"bytes,3,opt,name=metricName,proto3" json:"metricName,omitempty"`
}

func (x *ScaledObjectRef) Reset() {
	*x = ScaledObjectRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScaledObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScaledObjectRef) ProtoMessage() {}

func (x *ScaledObjectRef) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScaledObjectRef.ProtoReflect.Descriptor instead.
func (*ScaledObjectRef) Descriptor() ([]byte, []int) {
	return file_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *ScaledObjectRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScaledObjectRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ScaledObjectRef) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics     *v1beta1.ExternalMetricValueList `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
	PromMetrics *PromMetricsMsg                  `protobuf:"bytes,2,opt,name=promMetrics,proto3" json:"promMetrics,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMetrics() *v1beta1.ExternalMetricValueList {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *Response) GetPromMetrics() *PromMetricsMsg {
	if x != nil {
		return x.PromMetrics
	}
	return nil
}

// [DEPRECATED] PromMetricsMsg provides metrics for deprecated Prometheus Metrics in Metrics Server
type PromMetricsMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScaledObjectErr bool               `protobuf:"varint,1,opt,name=scaledObjectErr,proto3" json:"scaledObjectErr,omitempty"`
	ScalerMetric    []*ScalerMetricMsg `protobuf:"bytes,2,rep,name=scalerMetric,proto3" json:"scalerMetric,omitempty"`
	ScalerError     []*ScalerErrorMsg  `protobuf:"bytes,3,rep,name=scalerError,proto3" json:"scalerError,omitempty"`
}

func (x *PromMetricsMsg) Reset() {
	*x = PromMetricsMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromMetricsMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromMetricsMsg) ProtoMessage() {}

func (x *PromMetricsMsg) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromMetricsMsg.ProtoReflect.Descriptor instead.
func (*PromMetricsMsg) Descriptor() ([]byte, []int) {
	return file_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *PromMetricsMsg) GetScaledObjectErr() bool {
	if x != nil {
		return x.ScaledObjectErr
	}
	return false
}

func (x *PromMetricsMsg) GetScalerMetric() []*ScalerMetricMsg {
	if x != nil {
		return x.ScalerMetric
	}
	return nil
}

func (x *PromMetricsMsg) GetScalerError() []*ScalerErrorMsg {
	if x != nil {
		return x.ScalerError
	}
	return nil
}

type ScalerMetricMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScalerName  string  `protobuf:"bytes,1,opt,name=scalerName,proto3" json:"scalerName,omitempty"`
	ScalerIndex int32   `protobuf:"varint,2,opt,name=scalerIndex,proto3" json:"scalerIndex,omitempty"`
	MetricName  string  `protobuf:"bytes,3,opt,name=metricName,proto3" json:"metricName,omitempty"`
	MetricValue float32 `protobuf:"fixed32,4,opt,name=metricValue,proto3" json:"metricValue,omitempty"`
}

func (x *ScalerMetricMsg) Reset() {
	*x = ScalerMetricMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScalerMetricMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScalerMetricMsg) ProtoMessage() {}

func (x *ScalerMetricMsg) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScalerMetricMsg.ProtoReflect.Descriptor instead.
func (*ScalerMetricMsg) Descriptor() ([]byte, []int) {
	return file_metrics_proto_rawDescGZIP(), []int{3}
}

func (x *ScalerMetricMsg) GetScalerName() string {
	if x != nil {
		return x.ScalerName
	}
	return ""
}

func (x *ScalerMetricMsg) GetScalerIndex() int32 {
	if x != nil {
		return x.ScalerIndex
	}
	return 0
}

func (x *ScalerMetricMsg) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *ScalerMetricMsg) GetMetricValue() float32 {
	if x != nil {
		return x.MetricValue
	}
	return 0
}

type ScalerErrorMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScalerName  string `protobuf:"bytes,1,opt,name=scalerName,proto3" json:"scalerName,omitempty"`
	ScalerIndex int32  `protobuf:"varint,2,opt,name=scalerIndex,proto3" json:"scalerIndex,omitempty"`
	MetricName  string `protobuf:"bytes,3,opt,name=metricName,proto3" json:"metricName,omitempty"`
	Error       bool   `protobuf:"varint,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ScalerErrorMsg) Reset() {
	*x = ScalerErrorMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScalerErrorMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScalerErrorMsg) ProtoMessage() {}

func (x *ScalerErrorMsg) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScalerErrorMsg.ProtoReflect.Descriptor instead.
func (*ScalerErrorMsg) Descriptor() ([]byte, []int) {
	return file_metrics_proto_rawDescGZIP(), []int{4}
}

func (x *ScalerErrorMsg) GetScalerName() string {
	if x != nil {
		return x.ScalerName
	}
	return ""
}

func (x *ScalerErrorMsg) GetScalerIndex() int32 {
	if x != nil {
		return x.ScalerIndex
	}
	return 0
}

func (x *ScalerErrorMsg) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *ScalerErrorMsg) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_metrics_proto protoreflect.FileDescriptor

var file_metrics_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x1a, 0x40, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0f, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x64,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x6b, 0x38, 0x73, 0x2e,
	0x69, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x35, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x4d, 0x73, 0x67, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x72,
	0x72, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x63,
	0x61, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x73, 0x67, 0x52, 0x0c, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x35, 0x0a, 0x0b, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x53,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x45, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x63, 0x61, 0x6c,
	0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x1a, 0x0d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metrics_proto_rawDescOnce sync.Once
	file_metrics_proto_rawDescData = file_metrics_proto_rawDesc
)

func file_metrics_proto_rawDescGZIP() []byte {
	file_metrics_proto_rawDescOnce.Do(func() {
		file_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_metrics_proto_rawDescData)
	})
	return file_metrics_proto_rawDescData
}

var file_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_metrics_proto_goTypes = []interface{}{
	(*ScaledObjectRef)(nil),                 // 0: api.ScaledObjectRef
	(*Response)(nil),                        // 1: api.Response
	(*PromMetricsMsg)(nil),                  // 2: api.PromMetricsMsg
	(*ScalerMetricMsg)(nil),                 // 3: api.ScalerMetricMsg
	(*ScalerErrorMsg)(nil),                  // 4: api.ScalerErrorMsg
	(*v1beta1.ExternalMetricValueList)(nil), // 5: k8s.io.metrics.pkg.apis.external_metrics.v1beta1.ExternalMetricValueList
}
var file_metrics_proto_depIdxs = []int32{
	5, // 0: api.Response.metrics:type_name -> k8s.io.metrics.pkg.apis.external_metrics.v1beta1.ExternalMetricValueList
	2, // 1: api.Response.promMetrics:type_name -> api.PromMetricsMsg
	3, // 2: api.PromMetricsMsg.scalerMetric:type_name -> api.ScalerMetricMsg
	4, // 3: api.PromMetricsMsg.scalerError:type_name -> api.ScalerErrorMsg
	0, // 4: api.MetricsService.GetMetrics:input_type -> api.ScaledObjectRef
	1, // 5: api.MetricsService.GetMetrics:output_type -> api.Response
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_metrics_proto_init() }
func file_metrics_proto_init() {
	if File_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScaledObjectRef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromMetricsMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScalerMetricMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metrics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScalerErrorMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metrics_proto_goTypes,
		DependencyIndexes: file_metrics_proto_depIdxs,
		MessageInfos:      file_metrics_proto_msgTypes,
	}.Build()
	File_metrics_proto = out.File
	file_metrics_proto_rawDesc = nil
	file_metrics_proto_goTypes = nil
	file_metrics_proto_depIdxs = nil
}
