// Copyright 2016 The TensorFlow Authors. All Rights Reserved.
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
//==============================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: tensorflow/core/protobuf/master_service.proto

package for_core_protos_go_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_tensorflow_core_protobuf_master_service_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_master_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbb, 0x06, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x54, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x75, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x75, 0x6e,
	0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x12, 0x1a, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x65,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x4d, 0x61, 0x6b, 0x65, 0x43, 0x61,
	0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x52, 0x75, 0x6e,
	0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x8a, 0x01, 0x0a, 0x1a, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x42, 0x13, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tensorflow_core_protobuf_master_service_proto_goTypes = []interface{}{
	(*CreateSessionRequest)(nil),    // 0: tensorflow.CreateSessionRequest
	(*ExtendSessionRequest)(nil),    // 1: tensorflow.ExtendSessionRequest
	(*PartialRunSetupRequest)(nil),  // 2: tensorflow.PartialRunSetupRequest
	(*RunStepRequest)(nil),          // 3: tensorflow.RunStepRequest
	(*CloseSessionRequest)(nil),     // 4: tensorflow.CloseSessionRequest
	(*ListDevicesRequest)(nil),      // 5: tensorflow.ListDevicesRequest
	(*ResetRequest)(nil),            // 6: tensorflow.ResetRequest
	(*MakeCallableRequest)(nil),     // 7: tensorflow.MakeCallableRequest
	(*RunCallableRequest)(nil),      // 8: tensorflow.RunCallableRequest
	(*ReleaseCallableRequest)(nil),  // 9: tensorflow.ReleaseCallableRequest
	(*CreateSessionResponse)(nil),   // 10: tensorflow.CreateSessionResponse
	(*ExtendSessionResponse)(nil),   // 11: tensorflow.ExtendSessionResponse
	(*PartialRunSetupResponse)(nil), // 12: tensorflow.PartialRunSetupResponse
	(*RunStepResponse)(nil),         // 13: tensorflow.RunStepResponse
	(*CloseSessionResponse)(nil),    // 14: tensorflow.CloseSessionResponse
	(*ListDevicesResponse)(nil),     // 15: tensorflow.ListDevicesResponse
	(*ResetResponse)(nil),           // 16: tensorflow.ResetResponse
	(*MakeCallableResponse)(nil),    // 17: tensorflow.MakeCallableResponse
	(*RunCallableResponse)(nil),     // 18: tensorflow.RunCallableResponse
	(*ReleaseCallableResponse)(nil), // 19: tensorflow.ReleaseCallableResponse
}
var file_tensorflow_core_protobuf_master_service_proto_depIdxs = []int32{
	0,  // 0: tensorflow.grpc.MasterService.CreateSession:input_type -> tensorflow.CreateSessionRequest
	1,  // 1: tensorflow.grpc.MasterService.ExtendSession:input_type -> tensorflow.ExtendSessionRequest
	2,  // 2: tensorflow.grpc.MasterService.PartialRunSetup:input_type -> tensorflow.PartialRunSetupRequest
	3,  // 3: tensorflow.grpc.MasterService.RunStep:input_type -> tensorflow.RunStepRequest
	4,  // 4: tensorflow.grpc.MasterService.CloseSession:input_type -> tensorflow.CloseSessionRequest
	5,  // 5: tensorflow.grpc.MasterService.ListDevices:input_type -> tensorflow.ListDevicesRequest
	6,  // 6: tensorflow.grpc.MasterService.Reset:input_type -> tensorflow.ResetRequest
	7,  // 7: tensorflow.grpc.MasterService.MakeCallable:input_type -> tensorflow.MakeCallableRequest
	8,  // 8: tensorflow.grpc.MasterService.RunCallable:input_type -> tensorflow.RunCallableRequest
	9,  // 9: tensorflow.grpc.MasterService.ReleaseCallable:input_type -> tensorflow.ReleaseCallableRequest
	10, // 10: tensorflow.grpc.MasterService.CreateSession:output_type -> tensorflow.CreateSessionResponse
	11, // 11: tensorflow.grpc.MasterService.ExtendSession:output_type -> tensorflow.ExtendSessionResponse
	12, // 12: tensorflow.grpc.MasterService.PartialRunSetup:output_type -> tensorflow.PartialRunSetupResponse
	13, // 13: tensorflow.grpc.MasterService.RunStep:output_type -> tensorflow.RunStepResponse
	14, // 14: tensorflow.grpc.MasterService.CloseSession:output_type -> tensorflow.CloseSessionResponse
	15, // 15: tensorflow.grpc.MasterService.ListDevices:output_type -> tensorflow.ListDevicesResponse
	16, // 16: tensorflow.grpc.MasterService.Reset:output_type -> tensorflow.ResetResponse
	17, // 17: tensorflow.grpc.MasterService.MakeCallable:output_type -> tensorflow.MakeCallableResponse
	18, // 18: tensorflow.grpc.MasterService.RunCallable:output_type -> tensorflow.RunCallableResponse
	19, // 19: tensorflow.grpc.MasterService.ReleaseCallable:output_type -> tensorflow.ReleaseCallableResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_master_service_proto_init() }
func file_tensorflow_core_protobuf_master_service_proto_init() {
	if File_tensorflow_core_protobuf_master_service_proto != nil {
		return
	}
	file_tensorflow_core_protobuf_master_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_protobuf_master_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tensorflow_core_protobuf_master_service_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_master_service_proto_depIdxs,
	}.Build()
	File_tensorflow_core_protobuf_master_service_proto = out.File
	file_tensorflow_core_protobuf_master_service_proto_rawDesc = nil
	file_tensorflow_core_protobuf_master_service_proto_goTypes = nil
	file_tensorflow_core_protobuf_master_service_proto_depIdxs = nil
}
