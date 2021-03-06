/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

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

// Code generated by protoc-gen-go.
// source: github.com/intelsdi-x/snap/grpc/controlproxy/rpc/control.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	github.com/intelsdi-x/snap/grpc/controlproxy/rpc/control.proto

It has these top-level messages:
	SerrorReply
	PubProcMetricsRequest
	ErrorReply
	ProcessMetricsReply
	GetPluginContentTypesRequest
	GetPluginContentTypesReply
	ValidateDepsRequest
	ValidateDepsReply
	SubscribeDepsRequest
	SubscribeDepsReply
	Map
	MapEntry
	CollectMetricsRequest
	CollectMetricsReply
	ExpandWildcardsRequest
	ArrString
	ExpandWildcardsReply
	GetAutodiscoverPathsReply
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/intelsdi-x/snap/grpc/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SerrorReply struct {
}

func (m *SerrorReply) Reset()                    { *m = SerrorReply{} }
func (m *SerrorReply) String() string            { return proto.CompactTextString(m) }
func (*SerrorReply) ProtoMessage()               {}
func (*SerrorReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PubProcMetricsRequest struct {
	ContentType   string            `protobuf:"bytes,1,opt,name=ContentType,json=contentType" json:"ContentType,omitempty"`
	Content       []byte            `protobuf:"bytes,2,opt,name=Content,json=content,proto3" json:"Content,omitempty"`
	PluginName    string            `protobuf:"bytes,3,opt,name=PluginName,json=pluginName" json:"PluginName,omitempty"`
	PluginVersion int64             `protobuf:"varint,4,opt,name=PluginVersion,json=pluginVersion" json:"PluginVersion,omitempty"`
	Config        *common.ConfigMap `protobuf:"bytes,5,opt,name=Config,json=config" json:"Config,omitempty"`
	TaskId        string            `protobuf:"bytes,6,opt,name=TaskId,json=taskId" json:"TaskId,omitempty"`
}

func (m *PubProcMetricsRequest) Reset()                    { *m = PubProcMetricsRequest{} }
func (m *PubProcMetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*PubProcMetricsRequest) ProtoMessage()               {}
func (*PubProcMetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PubProcMetricsRequest) GetConfig() *common.ConfigMap {
	if m != nil {
		return m.Config
	}
	return nil
}

type ErrorReply struct {
	Errors []string `protobuf:"bytes,1,rep,name=Errors,json=errors" json:"Errors,omitempty"`
}

func (m *ErrorReply) Reset()                    { *m = ErrorReply{} }
func (m *ErrorReply) String() string            { return proto.CompactTextString(m) }
func (*ErrorReply) ProtoMessage()               {}
func (*ErrorReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ProcessMetricsReply struct {
	ContentType string   `protobuf:"bytes,1,opt,name=ContentType,json=contentType" json:"ContentType,omitempty"`
	Content     []byte   `protobuf:"bytes,2,opt,name=Content,json=content,proto3" json:"Content,omitempty"`
	Errors      []string `protobuf:"bytes,3,rep,name=Errors,json=errors" json:"Errors,omitempty"`
}

func (m *ProcessMetricsReply) Reset()                    { *m = ProcessMetricsReply{} }
func (m *ProcessMetricsReply) String() string            { return proto.CompactTextString(m) }
func (*ProcessMetricsReply) ProtoMessage()               {}
func (*ProcessMetricsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetPluginContentTypesRequest struct {
	Name       string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	PluginType int32  `protobuf:"varint,2,opt,name=PluginType,json=pluginType" json:"PluginType,omitempty"`
	Version    int32  `protobuf:"varint,3,opt,name=Version,json=version" json:"Version,omitempty"`
}

func (m *GetPluginContentTypesRequest) Reset()                    { *m = GetPluginContentTypesRequest{} }
func (m *GetPluginContentTypesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPluginContentTypesRequest) ProtoMessage()               {}
func (*GetPluginContentTypesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type GetPluginContentTypesReply struct {
	AcceptedTypes []string `protobuf:"bytes,1,rep,name=AcceptedTypes,json=acceptedTypes" json:"AcceptedTypes,omitempty"`
	ReturnedTypes []string `protobuf:"bytes,2,rep,name=ReturnedTypes,json=returnedTypes" json:"ReturnedTypes,omitempty"`
	Error         string   `protobuf:"bytes,3,opt,name=Error,json=error" json:"Error,omitempty"`
}

func (m *GetPluginContentTypesReply) Reset()                    { *m = GetPluginContentTypesReply{} }
func (m *GetPluginContentTypesReply) String() string            { return proto.CompactTextString(m) }
func (*GetPluginContentTypesReply) ProtoMessage()               {}
func (*GetPluginContentTypesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ValidateDepsRequest struct {
	Metrics []*common.Metric           `protobuf:"bytes,1,rep,name=Metrics,json=metrics" json:"Metrics,omitempty"`
	Plugins []*common.SubscribedPlugin `protobuf:"bytes,2,rep,name=Plugins,json=plugins" json:"Plugins,omitempty"`
}

func (m *ValidateDepsRequest) Reset()                    { *m = ValidateDepsRequest{} }
func (m *ValidateDepsRequest) String() string            { return proto.CompactTextString(m) }
func (*ValidateDepsRequest) ProtoMessage()               {}
func (*ValidateDepsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ValidateDepsRequest) GetMetrics() []*common.Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ValidateDepsRequest) GetPlugins() []*common.SubscribedPlugin {
	if m != nil {
		return m.Plugins
	}
	return nil
}

type ValidateDepsReply struct {
	Errors []*common.SnapError `protobuf:"bytes,1,rep,name=Errors,json=errors" json:"Errors,omitempty"`
}

func (m *ValidateDepsReply) Reset()                    { *m = ValidateDepsReply{} }
func (m *ValidateDepsReply) String() string            { return proto.CompactTextString(m) }
func (*ValidateDepsReply) ProtoMessage()               {}
func (*ValidateDepsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ValidateDepsReply) GetErrors() []*common.SnapError {
	if m != nil {
		return m.Errors
	}
	return nil
}

type SubscribeDepsRequest struct {
	Metrics []*common.Metric `protobuf:"bytes,1,rep,name=Metrics,json=metrics" json:"Metrics,omitempty"`
	Plugins []*common.Plugin `protobuf:"bytes,2,rep,name=Plugins,json=plugins" json:"Plugins,omitempty"`
	TaskId  string           `protobuf:"bytes,3,opt,name=TaskId,json=taskId" json:"TaskId,omitempty"`
}

func (m *SubscribeDepsRequest) Reset()                    { *m = SubscribeDepsRequest{} }
func (m *SubscribeDepsRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeDepsRequest) ProtoMessage()               {}
func (*SubscribeDepsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SubscribeDepsRequest) GetMetrics() []*common.Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *SubscribeDepsRequest) GetPlugins() []*common.Plugin {
	if m != nil {
		return m.Plugins
	}
	return nil
}

type SubscribeDepsReply struct {
	Errors []*common.SnapError `protobuf:"bytes,1,rep,name=Errors,json=errors" json:"Errors,omitempty"`
}

func (m *SubscribeDepsReply) Reset()                    { *m = SubscribeDepsReply{} }
func (m *SubscribeDepsReply) String() string            { return proto.CompactTextString(m) }
func (*SubscribeDepsReply) ProtoMessage()               {}
func (*SubscribeDepsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SubscribeDepsReply) GetErrors() []*common.SnapError {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Map struct {
	Entries []*MapEntry `protobuf:"bytes,1,rep,name=Entries,json=entries" json:"Entries,omitempty"`
}

func (m *Map) Reset()                    { *m = Map{} }
func (m *Map) String() string            { return proto.CompactTextString(m) }
func (*Map) ProtoMessage()               {}
func (*Map) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Map) GetEntries() []*MapEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type MapEntry struct {
	Key   string `protobuf:"bytes,1,opt,name=Key,json=key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,json=value" json:"Value,omitempty"`
}

func (m *MapEntry) Reset()                    { *m = MapEntry{} }
func (m *MapEntry) String() string            { return proto.CompactTextString(m) }
func (*MapEntry) ProtoMessage()               {}
func (*MapEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type CollectMetricsRequest struct {
	TaskID   string           `protobuf:"bytes,1,opt,name=TaskID,json=taskID" json:"TaskID,omitempty"`
	Metrics  []*common.Metric `protobuf:"bytes,2,rep,name=Metrics,json=metrics" json:"Metrics,omitempty"`
	Deadline *common.Time     `protobuf:"bytes,3,opt,name=Deadline,json=deadline" json:"Deadline,omitempty"`
	AllTags  map[string]*Map  `protobuf:"bytes,4,rep,name=AllTags,json=allTags" json:"AllTags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CollectMetricsRequest) Reset()                    { *m = CollectMetricsRequest{} }
func (m *CollectMetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*CollectMetricsRequest) ProtoMessage()               {}
func (*CollectMetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CollectMetricsRequest) GetMetrics() []*common.Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *CollectMetricsRequest) GetDeadline() *common.Time {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *CollectMetricsRequest) GetAllTags() map[string]*Map {
	if m != nil {
		return m.AllTags
	}
	return nil
}

type CollectMetricsReply struct {
	Metrics []*common.Metric `protobuf:"bytes,1,rep,name=Metrics,json=metrics" json:"Metrics,omitempty"`
	Errors  []string         `protobuf:"bytes,2,rep,name=Errors,json=errors" json:"Errors,omitempty"`
}

func (m *CollectMetricsReply) Reset()                    { *m = CollectMetricsReply{} }
func (m *CollectMetricsReply) String() string            { return proto.CompactTextString(m) }
func (*CollectMetricsReply) ProtoMessage()               {}
func (*CollectMetricsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *CollectMetricsReply) GetMetrics() []*common.Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type ExpandWildcardsRequest struct {
	Namespace []*common.NamespaceElement `protobuf:"bytes,1,rep,name=Namespace,json=namespace" json:"Namespace,omitempty"`
}

func (m *ExpandWildcardsRequest) Reset()                    { *m = ExpandWildcardsRequest{} }
func (m *ExpandWildcardsRequest) String() string            { return proto.CompactTextString(m) }
func (*ExpandWildcardsRequest) ProtoMessage()               {}
func (*ExpandWildcardsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ExpandWildcardsRequest) GetNamespace() []*common.NamespaceElement {
	if m != nil {
		return m.Namespace
	}
	return nil
}

type ArrString struct {
	S []*common.NamespaceElement `protobuf:"bytes,1,rep,name=S,json=s" json:"S,omitempty"`
}

func (m *ArrString) Reset()                    { *m = ArrString{} }
func (m *ArrString) String() string            { return proto.CompactTextString(m) }
func (*ArrString) ProtoMessage()               {}
func (*ArrString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ArrString) GetS() []*common.NamespaceElement {
	if m != nil {
		return m.S
	}
	return nil
}

type ExpandWildcardsReply struct {
	NSS   []*ArrString      `protobuf:"bytes,1,rep,name=NSS,json=nSS" json:"NSS,omitempty"`
	Error *common.SnapError `protobuf:"bytes,2,opt,name=Error,json=error" json:"Error,omitempty"`
}

func (m *ExpandWildcardsReply) Reset()                    { *m = ExpandWildcardsReply{} }
func (m *ExpandWildcardsReply) String() string            { return proto.CompactTextString(m) }
func (*ExpandWildcardsReply) ProtoMessage()               {}
func (*ExpandWildcardsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *ExpandWildcardsReply) GetNSS() []*ArrString {
	if m != nil {
		return m.NSS
	}
	return nil
}

func (m *ExpandWildcardsReply) GetError() *common.SnapError {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetAutodiscoverPathsReply struct {
	Paths []string `protobuf:"bytes,1,rep,name=Paths,json=paths" json:"Paths,omitempty"`
}

func (m *GetAutodiscoverPathsReply) Reset()                    { *m = GetAutodiscoverPathsReply{} }
func (m *GetAutodiscoverPathsReply) String() string            { return proto.CompactTextString(m) }
func (*GetAutodiscoverPathsReply) ProtoMessage()               {}
func (*GetAutodiscoverPathsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func init() {
	proto.RegisterType((*SerrorReply)(nil), "rpc.SerrorReply")
	proto.RegisterType((*PubProcMetricsRequest)(nil), "rpc.PubProcMetricsRequest")
	proto.RegisterType((*ErrorReply)(nil), "rpc.ErrorReply")
	proto.RegisterType((*ProcessMetricsReply)(nil), "rpc.ProcessMetricsReply")
	proto.RegisterType((*GetPluginContentTypesRequest)(nil), "rpc.GetPluginContentTypesRequest")
	proto.RegisterType((*GetPluginContentTypesReply)(nil), "rpc.GetPluginContentTypesReply")
	proto.RegisterType((*ValidateDepsRequest)(nil), "rpc.ValidateDepsRequest")
	proto.RegisterType((*ValidateDepsReply)(nil), "rpc.ValidateDepsReply")
	proto.RegisterType((*SubscribeDepsRequest)(nil), "rpc.SubscribeDepsRequest")
	proto.RegisterType((*SubscribeDepsReply)(nil), "rpc.SubscribeDepsReply")
	proto.RegisterType((*Map)(nil), "rpc.Map")
	proto.RegisterType((*MapEntry)(nil), "rpc.MapEntry")
	proto.RegisterType((*CollectMetricsRequest)(nil), "rpc.CollectMetricsRequest")
	proto.RegisterType((*CollectMetricsReply)(nil), "rpc.CollectMetricsReply")
	proto.RegisterType((*ExpandWildcardsRequest)(nil), "rpc.ExpandWildcardsRequest")
	proto.RegisterType((*ArrString)(nil), "rpc.ArrString")
	proto.RegisterType((*ExpandWildcardsReply)(nil), "rpc.ExpandWildcardsReply")
	proto.RegisterType((*GetAutodiscoverPathsReply)(nil), "rpc.GetAutodiscoverPathsReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for MetricManager service

type MetricManagerClient interface {
	// managesMetrics from scheduler
	GetPluginContentTypes(ctx context.Context, in *GetPluginContentTypesRequest, opts ...grpc.CallOption) (*GetPluginContentTypesReply, error)
	ExpandWildcards(ctx context.Context, in *ExpandWildcardsRequest, opts ...grpc.CallOption) (*ExpandWildcardsReply, error)
	CollectMetrics(ctx context.Context, in *CollectMetricsRequest, opts ...grpc.CallOption) (*CollectMetricsReply, error)
	PublishMetrics(ctx context.Context, in *PubProcMetricsRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	ProcessMetrics(ctx context.Context, in *PubProcMetricsRequest, opts ...grpc.CallOption) (*ProcessMetricsReply, error)
	ValidateDeps(ctx context.Context, in *ValidateDepsRequest, opts ...grpc.CallOption) (*ValidateDepsReply, error)
	SubscribeDeps(ctx context.Context, in *SubscribeDepsRequest, opts ...grpc.CallOption) (*SubscribeDepsReply, error)
	UnsubscribeDeps(ctx context.Context, in *SubscribeDepsRequest, opts ...grpc.CallOption) (*SubscribeDepsReply, error)
	MatchQueryToNamespaces(ctx context.Context, in *ExpandWildcardsRequest, opts ...grpc.CallOption) (*ExpandWildcardsReply, error)
	GetAutodiscoverPaths(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*GetAutodiscoverPathsReply, error)
}

type metricManagerClient struct {
	cc *grpc.ClientConn
}

func NewMetricManagerClient(cc *grpc.ClientConn) MetricManagerClient {
	return &metricManagerClient{cc}
}

func (c *metricManagerClient) GetPluginContentTypes(ctx context.Context, in *GetPluginContentTypesRequest, opts ...grpc.CallOption) (*GetPluginContentTypesReply, error) {
	out := new(GetPluginContentTypesReply)
	err := grpc.Invoke(ctx, "/rpc.MetricManager/GetPluginContentTypes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricManagerClient) ExpandWildcards(ctx context.Context, in *ExpandWildcardsRequest, opts ...grpc.CallOption) (*ExpandWildcardsReply, error) {
	out := new(ExpandWildcardsReply)
	err := grpc.Invoke(ctx, "/rpc.MetricManager/ExpandWildcards", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricManagerClient) CollectMetrics(ctx context.Context, in *CollectMetricsRequest, opts ...grpc.CallOption) (*CollectMetricsReply, error) {
	out := new(CollectMetricsReply)
	err := grpc.Invoke(ctx, "/rpc.MetricManager/CollectMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricManagerClient) PublishMetrics(ctx context.Context, in *PubProcMetricsRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := grpc.Invoke(ctx, "/rpc.MetricManager/PublishMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricManagerClient) ProcessMetrics(ctx context.Context, in *PubProcMetricsRequest, opts ...grpc.CallOption) (*ProcessMetricsReply, error) {
	out := new(ProcessMetricsReply)
	err := grpc.Invoke(ctx, "/rpc.MetricManager/ProcessMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricManagerClient) ValidateDeps(ctx context.Context, in *ValidateDepsRequest, opts ...grpc.CallOption) (*ValidateDepsReply, error) {
	out := new(ValidateDepsReply)
	err := grpc.Invoke(ctx, "/rpc.MetricManager/ValidateDeps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricManagerClient) SubscribeDeps(ctx context.Context, in *SubscribeDepsRequest, opts ...grpc.CallOption) (*SubscribeDepsReply, error) {
	out := new(SubscribeDepsReply)
	err := grpc.Invoke(ctx, "/rpc.MetricManager/SubscribeDeps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricManagerClient) UnsubscribeDeps(ctx context.Context, in *SubscribeDepsRequest, opts ...grpc.CallOption) (*SubscribeDepsReply, error) {
	out := new(SubscribeDepsReply)
	err := grpc.Invoke(ctx, "/rpc.MetricManager/UnsubscribeDeps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricManagerClient) MatchQueryToNamespaces(ctx context.Context, in *ExpandWildcardsRequest, opts ...grpc.CallOption) (*ExpandWildcardsReply, error) {
	out := new(ExpandWildcardsReply)
	err := grpc.Invoke(ctx, "/rpc.MetricManager/MatchQueryToNamespaces", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricManagerClient) GetAutodiscoverPaths(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*GetAutodiscoverPathsReply, error) {
	out := new(GetAutodiscoverPathsReply)
	err := grpc.Invoke(ctx, "/rpc.MetricManager/GetAutodiscoverPaths", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MetricManager service

type MetricManagerServer interface {
	// managesMetrics from scheduler
	GetPluginContentTypes(context.Context, *GetPluginContentTypesRequest) (*GetPluginContentTypesReply, error)
	ExpandWildcards(context.Context, *ExpandWildcardsRequest) (*ExpandWildcardsReply, error)
	CollectMetrics(context.Context, *CollectMetricsRequest) (*CollectMetricsReply, error)
	PublishMetrics(context.Context, *PubProcMetricsRequest) (*ErrorReply, error)
	ProcessMetrics(context.Context, *PubProcMetricsRequest) (*ProcessMetricsReply, error)
	ValidateDeps(context.Context, *ValidateDepsRequest) (*ValidateDepsReply, error)
	SubscribeDeps(context.Context, *SubscribeDepsRequest) (*SubscribeDepsReply, error)
	UnsubscribeDeps(context.Context, *SubscribeDepsRequest) (*SubscribeDepsReply, error)
	MatchQueryToNamespaces(context.Context, *ExpandWildcardsRequest) (*ExpandWildcardsReply, error)
	GetAutodiscoverPaths(context.Context, *common.Empty) (*GetAutodiscoverPathsReply, error)
}

func RegisterMetricManagerServer(s *grpc.Server, srv MetricManagerServer) {
	s.RegisterService(&_MetricManager_serviceDesc, srv)
}

func _MetricManager_GetPluginContentTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPluginContentTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricManagerServer).GetPluginContentTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MetricManager/GetPluginContentTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricManagerServer).GetPluginContentTypes(ctx, req.(*GetPluginContentTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricManager_ExpandWildcards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandWildcardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricManagerServer).ExpandWildcards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MetricManager/ExpandWildcards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricManagerServer).ExpandWildcards(ctx, req.(*ExpandWildcardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricManager_CollectMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricManagerServer).CollectMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MetricManager/CollectMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricManagerServer).CollectMetrics(ctx, req.(*CollectMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricManager_PublishMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubProcMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricManagerServer).PublishMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MetricManager/PublishMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricManagerServer).PublishMetrics(ctx, req.(*PubProcMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricManager_ProcessMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubProcMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricManagerServer).ProcessMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MetricManager/ProcessMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricManagerServer).ProcessMetrics(ctx, req.(*PubProcMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricManager_ValidateDeps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateDepsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricManagerServer).ValidateDeps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MetricManager/ValidateDeps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricManagerServer).ValidateDeps(ctx, req.(*ValidateDepsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricManager_SubscribeDeps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeDepsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricManagerServer).SubscribeDeps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MetricManager/SubscribeDeps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricManagerServer).SubscribeDeps(ctx, req.(*SubscribeDepsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricManager_UnsubscribeDeps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeDepsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricManagerServer).UnsubscribeDeps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MetricManager/UnsubscribeDeps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricManagerServer).UnsubscribeDeps(ctx, req.(*SubscribeDepsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricManager_MatchQueryToNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandWildcardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricManagerServer).MatchQueryToNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MetricManager/MatchQueryToNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricManagerServer).MatchQueryToNamespaces(ctx, req.(*ExpandWildcardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricManager_GetAutodiscoverPaths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricManagerServer).GetAutodiscoverPaths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.MetricManager/GetAutodiscoverPaths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricManagerServer).GetAutodiscoverPaths(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.MetricManager",
	HandlerType: (*MetricManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPluginContentTypes",
			Handler:    _MetricManager_GetPluginContentTypes_Handler,
		},
		{
			MethodName: "ExpandWildcards",
			Handler:    _MetricManager_ExpandWildcards_Handler,
		},
		{
			MethodName: "CollectMetrics",
			Handler:    _MetricManager_CollectMetrics_Handler,
		},
		{
			MethodName: "PublishMetrics",
			Handler:    _MetricManager_PublishMetrics_Handler,
		},
		{
			MethodName: "ProcessMetrics",
			Handler:    _MetricManager_ProcessMetrics_Handler,
		},
		{
			MethodName: "ValidateDeps",
			Handler:    _MetricManager_ValidateDeps_Handler,
		},
		{
			MethodName: "SubscribeDeps",
			Handler:    _MetricManager_SubscribeDeps_Handler,
		},
		{
			MethodName: "UnsubscribeDeps",
			Handler:    _MetricManager_UnsubscribeDeps_Handler,
		},
		{
			MethodName: "MatchQueryToNamespaces",
			Handler:    _MetricManager_MatchQueryToNamespaces_Handler,
		},
		{
			MethodName: "GetAutodiscoverPaths",
			Handler:    _MetricManager_GetAutodiscoverPaths_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() {
	proto.RegisterFile("github.com/intelsdi-x/snap/grpc/controlproxy/rpc/control.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 973 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0x4d, 0x6f, 0xdb, 0x46,
	0x13, 0x8e, 0x44, 0x4b, 0xb4, 0x47, 0x96, 0xfd, 0x66, 0xfd, 0xf1, 0x32, 0x6c, 0xe1, 0xa6, 0x44,
	0xd0, 0x38, 0x87, 0x4a, 0xa8, 0x0c, 0x14, 0x45, 0x0f, 0x09, 0xd4, 0x48, 0x70, 0x8b, 0xc0, 0x81,
	0x4a, 0xaa, 0xc9, 0xa9, 0x87, 0x15, 0xb9, 0x95, 0x09, 0x53, 0x24, 0xcb, 0x5d, 0x06, 0xd6, 0xa5,
	0x40, 0x7b, 0xee, 0x4f, 0xec, 0x8f, 0xe9, 0x7e, 0x91, 0x22, 0x55, 0xc6, 0x6e, 0x9b, 0x9e, 0xec,
	0x9d, 0x19, 0xce, 0xcc, 0xf3, 0xcc, 0xec, 0xb3, 0x82, 0xe7, 0xcb, 0x90, 0x5d, 0xe7, 0x8b, 0x81,
	0x9f, 0xac, 0x86, 0x61, 0xcc, 0x48, 0x44, 0x83, 0xf0, 0xf3, 0xdb, 0x21, 0x8d, 0x71, 0x3a, 0x5c,
	0x66, 0xa9, 0x3f, 0xf4, 0x93, 0x98, 0x65, 0x49, 0x94, 0x66, 0xc9, 0xed, 0x7a, 0x58, 0x31, 0x0c,
	0xb8, 0x85, 0x25, 0xc8, 0xe0, 0x26, 0xfb, 0xe2, 0xfe, 0x24, 0xab, 0x55, 0x12, 0xeb, 0x3f, 0xea,
	0x4b, 0xa7, 0x0f, 0x3d, 0x8f, 0x64, 0x59, 0x92, 0xb9, 0x24, 0x8d, 0xd6, 0xce, 0x1f, 0x2d, 0x38,
	0x99, 0xe5, 0x8b, 0x59, 0x96, 0xf8, 0x57, 0x84, 0x65, 0xa1, 0x4f, 0x5d, 0xf2, 0x73, 0x4e, 0x28,
	0x43, 0x8f, 0xa1, 0xf7, 0x92, 0xd7, 0x24, 0x31, 0x9b, 0xaf, 0x53, 0x62, 0xb5, 0x1e, 0xb7, 0xce,
	0xf7, 0xdc, 0x9e, 0xbf, 0x31, 0x21, 0x0b, 0x4c, 0x1d, 0x61, 0xb5, 0xb9, 0x77, 0xdf, 0x35, 0xb5,
	0x17, 0x9d, 0x01, 0xcc, 0xa2, 0x7c, 0x19, 0xc6, 0xaf, 0xf1, 0x8a, 0x58, 0x86, 0xfc, 0x14, 0xd2,
	0xd2, 0x82, 0x9e, 0x40, 0x5f, 0xf9, 0xdf, 0x90, 0x8c, 0x86, 0x49, 0x6c, 0xed, 0xf0, 0x10, 0xc3,
	0xed, 0xa7, 0x55, 0x23, 0x7a, 0x06, 0x5d, 0x9e, 0xff, 0xa7, 0x70, 0x69, 0x75, 0xb8, 0xbb, 0x37,
	0x7a, 0x38, 0xd0, 0x48, 0x94, 0xf5, 0x0a, 0xa7, 0x6e, 0xd7, 0x97, 0xff, 0xa2, 0x53, 0xe8, 0xce,
	0x31, 0xbd, 0xf9, 0x2e, 0xb0, 0xba, 0xb2, 0x58, 0x97, 0xc9, 0x93, 0xf3, 0x04, 0x60, 0x5a, 0x82,
	0x15, 0x51, 0xf2, 0x44, 0x39, 0x1a, 0x43, 0x44, 0x49, 0x22, 0xa8, 0x13, 0xc2, 0x91, 0x20, 0x80,
	0x50, 0x5a, 0x72, 0x20, 0xc2, 0x3f, 0x84, 0x81, 0x4d, 0x29, 0xa3, 0x56, 0x2a, 0x82, 0x8f, 0x2f,
	0x09, 0x53, 0xe0, 0x2b, 0xc9, 0x4b, 0xd6, 0x11, 0xec, 0x48, 0xce, 0x54, 0xb1, 0x9d, 0x58, 0xb0,
	0x55, 0xb2, 0x29, 0xdb, 0x10, 0x85, 0x3a, 0x05, 0x9b, 0x45, 0x17, 0x05, 0x8f, 0x86, 0x74, 0x9a,
	0xef, 0xd4, 0xd1, 0xf9, 0x05, 0xec, 0xf7, 0x54, 0x13, 0xf8, 0xf8, 0x14, 0xc6, 0xbe, 0x4f, 0x52,
	0x46, 0x02, 0x69, 0xd5, 0xac, 0xf4, 0x71, 0xd5, 0x28, 0xa2, 0x5c, 0xc2, 0xf2, 0x2c, 0x2e, 0xa2,
	0xda, 0x2a, 0x2a, 0xab, 0x1a, 0xd1, 0x31, 0x74, 0x24, 0x5e, 0x3d, 0xec, 0x8e, 0x84, 0xeb, 0x50,
	0x38, 0x7a, 0x83, 0xa3, 0x30, 0xc0, 0x8c, 0x4c, 0x48, 0x5a, 0x82, 0x3c, 0x07, 0x53, 0x13, 0x2d,
	0x4b, 0xf6, 0x46, 0x07, 0xc5, 0x64, 0x95, 0xd9, 0x35, 0x57, 0xca, 0x8d, 0x46, 0x60, 0xaa, 0xee,
	0x55, 0xd9, 0xde, 0xc8, 0x2a, 0x22, 0xbd, 0x7c, 0x41, 0xfd, 0x2c, 0x5c, 0x90, 0x40, 0x05, 0xb8,
	0xa6, 0x62, 0x84, 0x3a, 0xcf, 0xe1, 0x61, 0xbd, 0xa8, 0xc0, 0xfa, 0xac, 0x36, 0xfa, 0xca, 0x2e,
	0x79, 0xfc, 0xb6, 0xa8, 0x15, 0x29, 0x46, 0xf4, 0x5b, 0x0b, 0x8e, 0xcb, 0xec, 0xff, 0xae, 0xed,
	0xf3, 0xed, 0xb6, 0xcb, 0xc8, 0xad, 0x66, 0x2b, 0x8b, 0x6b, 0xd4, 0x16, 0xf7, 0x05, 0xa0, 0xad,
	0x1e, 0xfe, 0x21, 0x8a, 0x01, 0x18, 0xfc, 0x82, 0xa0, 0xa7, 0x60, 0x4e, 0xb9, 0x70, 0x84, 0xa4,
	0xf8, 0xa4, 0x3f, 0xe0, 0xca, 0x30, 0xe0, 0x2e, 0x61, 0x5e, 0xbb, 0x26, 0x51, 0x5e, 0x67, 0x04,
	0xbb, 0x85, 0x11, 0xfd, 0x0f, 0x8c, 0x57, 0x64, 0xad, 0x77, 0xd0, 0xb8, 0x21, 0x6b, 0x31, 0x5e,
	0xce, 0x69, 0xae, 0xb6, 0x8f, 0x8f, 0xf7, 0x9d, 0x38, 0x38, 0xbf, 0xb7, 0xe1, 0xe4, 0x65, 0x12,
	0x45, 0xc4, 0x67, 0x5b, 0xe2, 0x51, 0xc0, 0x9a, 0xe8, 0x24, 0x0a, 0xd6, 0xa4, 0x4a, 0x61, 0xfb,
	0x3e, 0x0a, 0x77, 0x27, 0x04, 0x07, 0x51, 0x18, 0x2b, 0x01, 0xe9, 0x8d, 0xf6, 0x8b, 0xd0, 0x79,
	0xb8, 0x22, 0xee, 0x6e, 0xa0, 0xbd, 0x68, 0x0c, 0xe6, 0x38, 0x8a, 0xe6, 0x78, 0x49, 0xb9, 0x8c,
	0x88, 0x9c, 0x4f, 0x25, 0xc4, 0xc6, 0xc6, 0x06, 0x3a, 0x52, 0x83, 0xc7, 0xea, 0x64, 0x4f, 0x60,
	0xbf, 0xea, 0x10, 0x04, 0xdc, 0xd4, 0x09, 0x38, 0x03, 0x85, 0x59, 0x12, 0xd0, 0x1b, 0xed, 0x16,
	0x2c, 0x6a, 0x2a, 0xbe, 0x6e, 0x7f, 0xd5, 0x72, 0xde, 0xc2, 0xd1, 0x76, 0x51, 0x31, 0xb4, 0xbf,
	0xbf, 0x36, 0x1b, 0xd1, 0x68, 0xd7, 0x44, 0x63, 0x06, 0xa7, 0xd3, 0xdb, 0x14, 0xc7, 0xc1, 0xdb,
	0x30, 0x0a, 0x7c, 0x9c, 0x05, 0x25, 0xcf, 0x5f, 0xc2, 0x9e, 0x90, 0x0b, 0x9a, 0x62, 0x9f, 0xe8,
	0xec, 0xe5, 0x0d, 0x29, 0x1d, 0xd3, 0x88, 0xac, 0xf8, 0x98, 0xdd, 0xbd, 0xb8, 0xb0, 0x38, 0x17,
	0xb0, 0x37, 0xce, 0x32, 0x8f, 0x97, 0x8d, 0x97, 0xe8, 0x33, 0x68, 0x79, 0xf7, 0x7e, 0xdc, 0xa2,
	0x0e, 0x86, 0xe3, 0xbf, 0xb4, 0xa1, 0x74, 0xd2, 0x78, 0xed, 0x79, 0x25, 0x38, 0xc1, 0x4c, 0x99,
	0xdc, 0x35, 0x62, 0xcf, 0xe3, 0x5b, 0xa8, 0xd5, 0xa1, 0x5d, 0x17, 0xf2, 0xcd, 0xda, 0x6a, 0xc1,
	0xf8, 0x02, 0x1e, 0x71, 0xc1, 0x1a, 0xe7, 0x2c, 0x09, 0x42, 0xea, 0x27, 0x5c, 0xc7, 0x66, 0x98,
	0x5d, 0xeb, 0x3a, 0x7c, 0x09, 0xe5, 0x49, 0xeb, 0x54, 0x27, 0x15, 0x87, 0xd1, 0xaf, 0x5d, 0xe8,
	0x2b, 0x22, 0xaf, 0x70, 0x8c, 0x97, 0x24, 0x43, 0x3f, 0xc2, 0x49, 0xa3, 0xea, 0xa1, 0x4f, 0x65,
	0x6f, 0x77, 0xe9, 0xaf, 0xfd, 0xc9, 0x5d, 0x21, 0xe2, 0xc1, 0x7c, 0x80, 0x5e, 0xc1, 0xe1, 0x16,
	0x0d, 0xe8, 0x23, 0xf9, 0x55, 0xf3, 0x8c, 0xec, 0x47, 0xcd, 0x4e, 0x95, 0xec, 0x5b, 0x38, 0xa8,
	0xef, 0x0c, 0xb2, 0xdf, 0xbf, 0xbd, 0xb6, 0xd5, 0xe8, 0x53, 0x99, 0x5e, 0xc0, 0x01, 0x7f, 0xc8,
	0xa3, 0x90, 0x5e, 0xd7, 0x33, 0x35, 0xbe, 0xee, 0xf6, 0xa1, 0x6a, 0x6a, 0xf3, 0x43, 0x40, 0xb6,
	0x52, 0x7f, 0x05, 0xef, 0x4c, 0xa0, 0x5a, 0x69, 0x78, 0x36, 0x79, 0xa6, 0x6f, 0x60, 0xbf, 0xaa,
	0xc0, 0x48, 0xc5, 0x36, 0xbc, 0x04, 0xf6, 0x69, 0x83, 0x47, 0xe5, 0x98, 0x42, 0xbf, 0x26, 0x80,
	0x48, 0xd1, 0xd8, 0x24, 0xcc, 0xf6, 0xff, 0x9b, 0x5c, 0x2a, 0xcd, 0x25, 0x1c, 0xfe, 0x10, 0xd3,
	0xff, 0x20, 0x91, 0x0b, 0xa7, 0x57, 0x98, 0xf9, 0xd7, 0xdf, 0xe7, 0x24, 0x5b, 0xcf, 0x93, 0xf2,
	0x7e, 0x7c, 0xc8, 0xf0, 0x2f, 0xe1, 0xb8, 0x69, 0xdb, 0x51, 0xbf, 0xb8, 0x1f, 0xd3, 0x55, 0xca,
	0xd6, 0xf6, 0x59, 0xb1, 0x93, 0xcd, 0xf7, 0xc2, 0x79, 0xb0, 0xe8, 0xca, 0xdf, 0x76, 0x17, 0x7f,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xca, 0xcd, 0x72, 0x57, 0x0a, 0x00, 0x00,
}
