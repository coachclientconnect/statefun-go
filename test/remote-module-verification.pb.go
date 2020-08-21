//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: src/main/protobuf/remote-module-verification.proto

package main

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Invoke struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Invoke) Reset() {
	*x = Invoke{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_main_protobuf_remote_module_verification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoke) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoke) ProtoMessage() {}

func (x *Invoke) ProtoReflect() protoreflect.Message {
	mi := &file_src_main_protobuf_remote_module_verification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoke.ProtoReflect.Descriptor instead.
func (*Invoke) Descriptor() ([]byte, []int) {
	return file_src_main_protobuf_remote_module_verification_proto_rawDescGZIP(), []int{0}
}

type InvokeCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *InvokeCount) Reset() {
	*x = InvokeCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_main_protobuf_remote_module_verification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeCount) ProtoMessage() {}

func (x *InvokeCount) ProtoReflect() protoreflect.Message {
	mi := &file_src_main_protobuf_remote_module_verification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeCount.ProtoReflect.Descriptor instead.
func (*InvokeCount) Descriptor() ([]byte, []int) {
	return file_src_main_protobuf_remote_module_verification_proto_rawDescGZIP(), []int{1}
}

func (x *InvokeCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type InvokeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InvokeCount int32  `protobuf:"varint,2,opt,name=invoke_count,json=invokeCount,proto3" json:"invoke_count,omitempty"`
}

func (x *InvokeResult) Reset() {
	*x = InvokeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_main_protobuf_remote_module_verification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeResult) ProtoMessage() {}

func (x *InvokeResult) ProtoReflect() protoreflect.Message {
	mi := &file_src_main_protobuf_remote_module_verification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeResult.ProtoReflect.Descriptor instead.
func (*InvokeResult) Descriptor() ([]byte, []int) {
	return file_src_main_protobuf_remote_module_verification_proto_rawDescGZIP(), []int{2}
}

func (x *InvokeResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InvokeResult) GetInvokeCount() int32 {
	if x != nil {
		return x.InvokeCount
	}
	return 0
}

var File_src_main_protobuf_remote_module_verification_proto protoreflect.FileDescriptor

var file_src_main_protobuf_remote_module_verification_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e,
	0x65, 0x32, 0x65, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x22, 0x08, 0x0a, 0x06, 0x49, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x0c, 0x49, 0x6e, 0x76,
	0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76,
	0x6f, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x38, 0x0a, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x50, 0x00,
	0x5a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_main_protobuf_remote_module_verification_proto_rawDescOnce sync.Once
	file_src_main_protobuf_remote_module_verification_proto_rawDescData = file_src_main_protobuf_remote_module_verification_proto_rawDesc
)

func file_src_main_protobuf_remote_module_verification_proto_rawDescGZIP() []byte {
	file_src_main_protobuf_remote_module_verification_proto_rawDescOnce.Do(func() {
		file_src_main_protobuf_remote_module_verification_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_main_protobuf_remote_module_verification_proto_rawDescData)
	})
	return file_src_main_protobuf_remote_module_verification_proto_rawDescData
}

var file_src_main_protobuf_remote_module_verification_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_main_protobuf_remote_module_verification_proto_goTypes = []interface{}{
	(*Invoke)(nil),       // 0: org.apache.flink.statefun.e2e.remote.Invoke
	(*InvokeCount)(nil),  // 1: org.apache.flink.statefun.e2e.remote.InvokeCount
	(*InvokeResult)(nil), // 2: org.apache.flink.statefun.e2e.remote.InvokeResult
}
var file_src_main_protobuf_remote_module_verification_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_main_protobuf_remote_module_verification_proto_init() }
func file_src_main_protobuf_remote_module_verification_proto_init() {
	if File_src_main_protobuf_remote_module_verification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_main_protobuf_remote_module_verification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoke); i {
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
		file_src_main_protobuf_remote_module_verification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeCount); i {
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
		file_src_main_protobuf_remote_module_verification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeResult); i {
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
			RawDescriptor: file_src_main_protobuf_remote_module_verification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_main_protobuf_remote_module_verification_proto_goTypes,
		DependencyIndexes: file_src_main_protobuf_remote_module_verification_proto_depIdxs,
		MessageInfos:      file_src_main_protobuf_remote_module_verification_proto_msgTypes,
	}.Build()
	File_src_main_protobuf_remote_module_verification_proto = out.File
	file_src_main_protobuf_remote_module_verification_proto_rawDesc = nil
	file_src_main_protobuf_remote_module_verification_proto_goTypes = nil
	file_src_main_protobuf_remote_module_verification_proto_depIdxs = nil
}