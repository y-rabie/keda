// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/cloud/nexus/v1/message.proto

package nexus

import (
	reflect "reflect"
	sync "sync"

	v1 "go.temporal.io/api/cloud/resource/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EndpointSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the endpoint. Must be unique within an account.
	// The name must match `^[a-zA-Z][a-zA-Z0-9\-]*[a-zA-Z0-9]$`.
	// This field is mutable.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Indicates where the endpoint should forward received nexus requests to.
	TargetSpec *EndpointTargetSpec `protobuf:"bytes,2,opt,name=target_spec,json=targetSpec,proto3" json:"target_spec,omitempty"`
	// The set of policies (e.g. authorization) for the endpoint. Each request's caller
	// must match with at least one of the specs to be accepted by the endpoint.
	// This field is mutable.
	PolicySpecs []*EndpointPolicySpec `protobuf:"bytes,3,rep,name=policy_specs,json=policySpecs,proto3" json:"policy_specs,omitempty"`
	// The markdown description of the endpoint - optional.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EndpointSpec) Reset() {
	*x = EndpointSpec{}
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndpointSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointSpec) ProtoMessage() {}

func (x *EndpointSpec) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointSpec.ProtoReflect.Descriptor instead.
func (*EndpointSpec) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_nexus_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *EndpointSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EndpointSpec) GetTargetSpec() *EndpointTargetSpec {
	if x != nil {
		return x.TargetSpec
	}
	return nil
}

func (x *EndpointSpec) GetPolicySpecs() []*EndpointPolicySpec {
	if x != nil {
		return x.PolicySpecs
	}
	return nil
}

func (x *EndpointSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type EndpointTargetSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Variant:
	//
	//	*EndpointTargetSpec_WorkerTargetSpec
	Variant isEndpointTargetSpec_Variant `protobuf_oneof:"variant"`
}

func (x *EndpointTargetSpec) Reset() {
	*x = EndpointTargetSpec{}
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndpointTargetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointTargetSpec) ProtoMessage() {}

func (x *EndpointTargetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointTargetSpec.ProtoReflect.Descriptor instead.
func (*EndpointTargetSpec) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_nexus_v1_message_proto_rawDescGZIP(), []int{1}
}

func (m *EndpointTargetSpec) GetVariant() isEndpointTargetSpec_Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (x *EndpointTargetSpec) GetWorkerTargetSpec() *WorkerTargetSpec {
	if x, ok := x.GetVariant().(*EndpointTargetSpec_WorkerTargetSpec); ok {
		return x.WorkerTargetSpec
	}
	return nil
}

type isEndpointTargetSpec_Variant interface {
	isEndpointTargetSpec_Variant()
}

type EndpointTargetSpec_WorkerTargetSpec struct {
	// A target spec for routing nexus requests to a specific cloud namespace worker.
	WorkerTargetSpec *WorkerTargetSpec `protobuf:"bytes,1,opt,name=worker_target_spec,json=workerTargetSpec,proto3,oneof"`
}

func (*EndpointTargetSpec_WorkerTargetSpec) isEndpointTargetSpec_Variant() {}

type WorkerTargetSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The target cloud namespace to route requests to. Namespace must be in same account as the endpoint. This field is mutable.
	NamespaceId string `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	// The task queue on the cloud namespace to route requests to. This field is mutable.
	TaskQueue string `protobuf:"bytes,2,opt,name=task_queue,json=taskQueue,proto3" json:"task_queue,omitempty"`
}

func (x *WorkerTargetSpec) Reset() {
	*x = WorkerTargetSpec{}
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkerTargetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerTargetSpec) ProtoMessage() {}

func (x *WorkerTargetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerTargetSpec.ProtoReflect.Descriptor instead.
func (*WorkerTargetSpec) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_nexus_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *WorkerTargetSpec) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *WorkerTargetSpec) GetTaskQueue() string {
	if x != nil {
		return x.TaskQueue
	}
	return ""
}

type EndpointPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Variant:
	//
	//	*EndpointPolicySpec_AllowedCloudNamespacePolicySpec
	Variant isEndpointPolicySpec_Variant `protobuf_oneof:"variant"`
}

func (x *EndpointPolicySpec) Reset() {
	*x = EndpointPolicySpec{}
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndpointPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointPolicySpec) ProtoMessage() {}

func (x *EndpointPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointPolicySpec.ProtoReflect.Descriptor instead.
func (*EndpointPolicySpec) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_nexus_v1_message_proto_rawDescGZIP(), []int{3}
}

func (m *EndpointPolicySpec) GetVariant() isEndpointPolicySpec_Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (x *EndpointPolicySpec) GetAllowedCloudNamespacePolicySpec() *AllowedCloudNamespacePolicySpec {
	if x, ok := x.GetVariant().(*EndpointPolicySpec_AllowedCloudNamespacePolicySpec); ok {
		return x.AllowedCloudNamespacePolicySpec
	}
	return nil
}

type isEndpointPolicySpec_Variant interface {
	isEndpointPolicySpec_Variant()
}

type EndpointPolicySpec_AllowedCloudNamespacePolicySpec struct {
	// A policy spec that allows one caller namespace to access the endpoint.
	AllowedCloudNamespacePolicySpec *AllowedCloudNamespacePolicySpec `protobuf:"bytes,1,opt,name=allowed_cloud_namespace_policy_spec,json=allowedCloudNamespacePolicySpec,proto3,oneof"`
}

func (*EndpointPolicySpec_AllowedCloudNamespacePolicySpec) isEndpointPolicySpec_Variant() {}

type AllowedCloudNamespacePolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace that is allowed to call into this endpoint. Calling namespace must be in same account as the endpoint.
	NamespaceId string `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
}

func (x *AllowedCloudNamespacePolicySpec) Reset() {
	*x = AllowedCloudNamespacePolicySpec{}
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllowedCloudNamespacePolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedCloudNamespacePolicySpec) ProtoMessage() {}

func (x *AllowedCloudNamespacePolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowedCloudNamespacePolicySpec.ProtoReflect.Descriptor instead.
func (*AllowedCloudNamespacePolicySpec) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_nexus_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *AllowedCloudNamespacePolicySpec) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

// An endpoint that receives and then routes Nexus requests
type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the endpoint. This is generated by the server and is immutable.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The current version of the endpoint specification.
	// The next update operation must include this version.
	ResourceVersion string `protobuf:"bytes,2,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	// The endpoint specification.
	Spec *EndpointSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	// The current state of the endpoint.
	// For any failed state, reach out to Temporal Cloud support for remediation.
	State v1.ResourceState `protobuf:"varint,4,opt,name=state,proto3,enum=temporal.api.cloud.resource.v1.ResourceState" json:"state,omitempty"`
	// The id of any ongoing async operation that is creating, updating, or deleting the endpoint, if any.
	AsyncOperationId string `protobuf:"bytes,5,opt,name=async_operation_id,json=asyncOperationId,proto3" json:"async_operation_id,omitempty"`
	// The date and time when the endpoint was created.
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// The date and time when the endpoint was last modified.
	LastModifiedTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_modified_time,json=lastModifiedTime,proto3" json:"last_modified_time,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_nexus_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *Endpoint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Endpoint) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *Endpoint) GetSpec() *EndpointSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Endpoint) GetState() v1.ResourceState {
	if x != nil {
		return x.State
	}
	return v1.ResourceState(0)
}

func (x *Endpoint) GetAsyncOperationId() string {
	if x != nil {
		return x.AsyncOperationId
	}
	return ""
}

func (x *Endpoint) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Endpoint) GetLastModifiedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModifiedTime
	}
	return nil
}

var File_temporal_api_cloud_nexus_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_cloud_nexus_v1_message_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x78, 0x75,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfa, 0x01, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x16, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x54, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e,
	0x65, 0x78, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x42, 0x02, 0x68, 0x00, 0x12, 0x56, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0b, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x24, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00,
	0x22, 0x82, 0x01, 0x0a, 0x12, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x61, 0x0a, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52,
	0x10, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x42, 0x02, 0x68, 0x00, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x22,
	0x5c, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x25, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x21, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x42, 0x02, 0x68, 0x00, 0x22, 0xb2, 0x01, 0x0a, 0x12, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x90,
	0x01, 0x0a, 0x23, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e,
	0x65, 0x78, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x1f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x42, 0x02, 0x68, 0x00, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x1f, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x25, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x22, 0x9c, 0x03, 0x0a, 0x08,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x2d, 0x0a, 0x10, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x12, 0x41, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x42, 0x02, 0x68, 0x00, 0x12, 0x47, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x30, 0x0a, 0x12,
	0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x41, 0x0a, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4c, 0x0a, 0x12, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_cloud_nexus_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_cloud_nexus_v1_message_proto_rawDescData = file_temporal_api_cloud_nexus_v1_message_proto_rawDesc
)

func file_temporal_api_cloud_nexus_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_cloud_nexus_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_cloud_nexus_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_cloud_nexus_v1_message_proto_rawDescData)
	})
	return file_temporal_api_cloud_nexus_v1_message_proto_rawDescData
}

var file_temporal_api_cloud_nexus_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_temporal_api_cloud_nexus_v1_message_proto_goTypes = []any{
	(*EndpointSpec)(nil),                    // 0: temporal.api.cloud.nexus.v1.EndpointSpec
	(*EndpointTargetSpec)(nil),              // 1: temporal.api.cloud.nexus.v1.EndpointTargetSpec
	(*WorkerTargetSpec)(nil),                // 2: temporal.api.cloud.nexus.v1.WorkerTargetSpec
	(*EndpointPolicySpec)(nil),              // 3: temporal.api.cloud.nexus.v1.EndpointPolicySpec
	(*AllowedCloudNamespacePolicySpec)(nil), // 4: temporal.api.cloud.nexus.v1.AllowedCloudNamespacePolicySpec
	(*Endpoint)(nil),                        // 5: temporal.api.cloud.nexus.v1.Endpoint
	(v1.ResourceState)(0),                   // 6: temporal.api.cloud.resource.v1.ResourceState
	(*timestamppb.Timestamp)(nil),           // 7: google.protobuf.Timestamp
}
var file_temporal_api_cloud_nexus_v1_message_proto_depIdxs = []int32{
	1, // 0: temporal.api.cloud.nexus.v1.EndpointSpec.target_spec:type_name -> temporal.api.cloud.nexus.v1.EndpointTargetSpec
	3, // 1: temporal.api.cloud.nexus.v1.EndpointSpec.policy_specs:type_name -> temporal.api.cloud.nexus.v1.EndpointPolicySpec
	2, // 2: temporal.api.cloud.nexus.v1.EndpointTargetSpec.worker_target_spec:type_name -> temporal.api.cloud.nexus.v1.WorkerTargetSpec
	4, // 3: temporal.api.cloud.nexus.v1.EndpointPolicySpec.allowed_cloud_namespace_policy_spec:type_name -> temporal.api.cloud.nexus.v1.AllowedCloudNamespacePolicySpec
	0, // 4: temporal.api.cloud.nexus.v1.Endpoint.spec:type_name -> temporal.api.cloud.nexus.v1.EndpointSpec
	6, // 5: temporal.api.cloud.nexus.v1.Endpoint.state:type_name -> temporal.api.cloud.resource.v1.ResourceState
	7, // 6: temporal.api.cloud.nexus.v1.Endpoint.created_time:type_name -> google.protobuf.Timestamp
	7, // 7: temporal.api.cloud.nexus.v1.Endpoint.last_modified_time:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_temporal_api_cloud_nexus_v1_message_proto_init() }
func file_temporal_api_cloud_nexus_v1_message_proto_init() {
	if File_temporal_api_cloud_nexus_v1_message_proto != nil {
		return
	}
	file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[1].OneofWrappers = []any{
		(*EndpointTargetSpec_WorkerTargetSpec)(nil),
	}
	file_temporal_api_cloud_nexus_v1_message_proto_msgTypes[3].OneofWrappers = []any{
		(*EndpointPolicySpec_AllowedCloudNamespacePolicySpec)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_cloud_nexus_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_cloud_nexus_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_cloud_nexus_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_cloud_nexus_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_cloud_nexus_v1_message_proto = out.File
	file_temporal_api_cloud_nexus_v1_message_proto_rawDesc = nil
	file_temporal_api_cloud_nexus_v1_message_proto_goTypes = nil
	file_temporal_api_cloud_nexus_v1_message_proto_depIdxs = nil
}
