// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: kuscia/proto/api/v1alpha1/kusciaapi/model.proto

package kusciaapi

import (
	v1alpha1 "github.com/secretflow/kuscia/proto/api/v1alpha1"
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

type UploadModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// model filename
	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	// model file content(base64)
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UploadModelRequest) Reset() {
	*x = UploadModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadModelRequest) ProtoMessage() {}

func (x *UploadModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadModelRequest.ProtoReflect.Descriptor instead.
func (*UploadModelRequest) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDescGZIP(), []int{0}
}

func (x *UploadModelRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadModelRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UploadModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *v1alpha1.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UploadModelResponse) Reset() {
	*x = UploadModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadModelResponse) ProtoMessage() {}

func (x *UploadModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadModelResponse.ProtoReflect.Descriptor instead.
func (*UploadModelResponse) Descriptor() ([]byte, []int) {
	return file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDescGZIP(), []int{1}
}

func (x *UploadModelResponse) GetStatus() *v1alpha1.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_kuscia_proto_api_v1alpha1_kusciaapi_model_proto protoreflect.FileDescriptor

var file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b, 0x75, 0x73, 0x63,
	0x69, 0x61, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x23, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6b, 0x75, 0x73,
	0x63, 0x69, 0x61, 0x61, 0x70, 0x69, 0x1a, 0x26, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a,
	0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x13, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x91, 0x01, 0x0a,
	0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01,
	0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x37, 0x2e,
	0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x5e, 0x0a, 0x21, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x6b, 0x75, 0x73, 0x63,
	0x69, 0x61, 0x61, 0x70, 0x69, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6b, 0x75, 0x73,
	0x63, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDescOnce sync.Once
	file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDescData = file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDesc
)

func file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDescGZIP() []byte {
	file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDescOnce.Do(func() {
		file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDescData)
	})
	return file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDescData
}

var file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_goTypes = []interface{}{
	(*UploadModelRequest)(nil),  // 0: kuscia.proto.api.v1alpha1.kusciaapi.UploadModelRequest
	(*UploadModelResponse)(nil), // 1: kuscia.proto.api.v1alpha1.kusciaapi.UploadModelResponse
	(*v1alpha1.Status)(nil),     // 2: kuscia.proto.api.v1alpha1.Status
}
var file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_depIdxs = []int32{
	2, // 0: kuscia.proto.api.v1alpha1.kusciaapi.UploadModelResponse.status:type_name -> kuscia.proto.api.v1alpha1.Status
	0, // 1: kuscia.proto.api.v1alpha1.kusciaapi.ModelService.UploadModel:input_type -> kuscia.proto.api.v1alpha1.kusciaapi.UploadModelRequest
	1, // 2: kuscia.proto.api.v1alpha1.kusciaapi.ModelService.UploadModel:output_type -> kuscia.proto.api.v1alpha1.kusciaapi.UploadModelResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_init() }
func file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_init() {
	if File_kuscia_proto_api_v1alpha1_kusciaapi_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadModelRequest); i {
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
		file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadModelResponse); i {
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
			RawDescriptor: file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_goTypes,
		DependencyIndexes: file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_depIdxs,
		MessageInfos:      file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_msgTypes,
	}.Build()
	File_kuscia_proto_api_v1alpha1_kusciaapi_model_proto = out.File
	file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_rawDesc = nil
	file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_goTypes = nil
	file_kuscia_proto_api_v1alpha1_kusciaapi_model_proto_depIdxs = nil
}
