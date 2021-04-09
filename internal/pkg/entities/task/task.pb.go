// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: internal/pkg/entities/task/task.proto

package task

import (
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

type Status int32

const (
	Status_BACKLOG    Status = 0
	Status_TODO       Status = 1
	Status_INPROGRESS Status = 2
	Status_DONE       Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "BACKLOG",
		1: "TODO",
		2: "INPROGRESS",
		3: "DONE",
	}
	Status_value = map[string]int32{
		"BACKLOG":    0,
		"TODO":       1,
		"INPROGRESS": 2,
		"DONE":       3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_pkg_entities_task_task_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_internal_pkg_entities_task_task_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_internal_pkg_entities_task_task_proto_rawDescGZIP(), []int{0}
}

var File_internal_pkg_entities_task_task_proto protoreflect.FileDescriptor

var file_internal_pkg_entities_task_task_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x2a, 0x39, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x41, 0x43, 0x4b, 0x4c,
	0x4f, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x44, 0x4f, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x49, 0x4e, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_entities_task_task_proto_rawDescOnce sync.Once
	file_internal_pkg_entities_task_task_proto_rawDescData = file_internal_pkg_entities_task_task_proto_rawDesc
)

func file_internal_pkg_entities_task_task_proto_rawDescGZIP() []byte {
	file_internal_pkg_entities_task_task_proto_rawDescOnce.Do(func() {
		file_internal_pkg_entities_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_entities_task_task_proto_rawDescData)
	})
	return file_internal_pkg_entities_task_task_proto_rawDescData
}

var file_internal_pkg_entities_task_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_pkg_entities_task_task_proto_goTypes = []interface{}{
	(Status)(0), // 0: task.Status
}
var file_internal_pkg_entities_task_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pkg_entities_task_task_proto_init() }
func file_internal_pkg_entities_task_task_proto_init() {
	if File_internal_pkg_entities_task_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pkg_entities_task_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pkg_entities_task_task_proto_goTypes,
		DependencyIndexes: file_internal_pkg_entities_task_task_proto_depIdxs,
		EnumInfos:         file_internal_pkg_entities_task_task_proto_enumTypes,
	}.Build()
	File_internal_pkg_entities_task_task_proto = out.File
	file_internal_pkg_entities_task_task_proto_rawDesc = nil
	file_internal_pkg_entities_task_task_proto_goTypes = nil
	file_internal_pkg_entities_task_task_proto_depIdxs = nil
}
