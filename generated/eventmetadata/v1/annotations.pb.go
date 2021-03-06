// Copyright 2022 Dhi Aurrahman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: eventmetadata/v1/annotations.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_eventmetadata_v1_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*EventMetadata)(nil),
		Field:         45081702,
		Name:          "eventmetadata.v1.event_metadata",
		Tag:           "bytes,45081702,opt,name=event_metadata",
		Filename:      "eventmetadata/v1/annotations.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional eventmetadata.v1.EventMetadata event_metadata = 45081702;
	E_EventMetadata = &file_eventmetadata_v1_annotations_proto_extTypes[0]
)

var File_eventmetadata_v1_annotations_proto protoreflect.FileDescriptor

var file_eventmetadata_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x6a, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe6, 0xc8, 0xbf, 0x15, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x69, 0x6f, 0x2f, 0x6b, 0x61, 0x6b, 0x61, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_eventmetadata_v1_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*EventMetadata)(nil),               // 1: eventmetadata.v1.EventMetadata
}
var file_eventmetadata_v1_annotations_proto_depIdxs = []int32{
	0, // 0: eventmetadata.v1.event_metadata:extendee -> google.protobuf.MessageOptions
	1, // 1: eventmetadata.v1.event_metadata:type_name -> eventmetadata.v1.EventMetadata
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eventmetadata_v1_annotations_proto_init() }
func file_eventmetadata_v1_annotations_proto_init() {
	if File_eventmetadata_v1_annotations_proto != nil {
		return
	}
	file_eventmetadata_v1_schema_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eventmetadata_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_eventmetadata_v1_annotations_proto_goTypes,
		DependencyIndexes: file_eventmetadata_v1_annotations_proto_depIdxs,
		ExtensionInfos:    file_eventmetadata_v1_annotations_proto_extTypes,
	}.Build()
	File_eventmetadata_v1_annotations_proto = out.File
	file_eventmetadata_v1_annotations_proto_rawDesc = nil
	file_eventmetadata_v1_annotations_proto_goTypes = nil
	file_eventmetadata_v1_annotations_proto_depIdxs = nil
}
