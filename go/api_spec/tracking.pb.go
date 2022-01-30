// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.8
// source: tracking.proto

package api_spec

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AkitaWitnessTracking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Earliest recorded occurrence, inclusive.
	FirstSeen *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=first_seen,json=firstSeen,proto3" json:"first_seen,omitempty"`
	// Most recent recorded occurrence, inclusive, as an offset in seconds
	// from first_seen.
	LastSeenOffsetSeconds uint32 `protobuf:"varint,2,opt,name=last_seen_offset_seconds,json=lastSeenOffsetSeconds,proto3" json:"last_seen_offset_seconds,omitempty"`
	// Total number of occurrences.
	Count uint64 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AkitaWitnessTracking) Reset() {
	*x = AkitaWitnessTracking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AkitaWitnessTracking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AkitaWitnessTracking) ProtoMessage() {}

func (x *AkitaWitnessTracking) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AkitaWitnessTracking.ProtoReflect.Descriptor instead.
func (*AkitaWitnessTracking) Descriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{0}
}

func (x *AkitaWitnessTracking) GetFirstSeen() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstSeen
	}
	return nil
}

func (x *AkitaWitnessTracking) GetLastSeenOffsetSeconds() uint32 {
	if x != nil {
		return x.LastSeenOffsetSeconds
	}
	return 0
}

func (x *AkitaWitnessTracking) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_tracking_proto protoreflect.FileDescriptor

var file_tracking_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x14,
	0x41, 0x6b, 0x69, 0x74, 0x61, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x12,
	0x37, 0x0a, 0x18, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x2c,
	0x5a, 0x2a, 0x61, 0x6b, 0x69, 0x74, 0x61, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x73, 0x74, 0x61, 0x72, 0x2f, 0x70, 0x62,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracking_proto_rawDescOnce sync.Once
	file_tracking_proto_rawDescData = file_tracking_proto_rawDesc
)

func file_tracking_proto_rawDescGZIP() []byte {
	file_tracking_proto_rawDescOnce.Do(func() {
		file_tracking_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracking_proto_rawDescData)
	})
	return file_tracking_proto_rawDescData
}

var file_tracking_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tracking_proto_goTypes = []interface{}{
	(*AkitaWitnessTracking)(nil),  // 0: api_spec.AkitaWitnessTracking
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_tracking_proto_depIdxs = []int32{
	1, // 0: api_spec.AkitaWitnessTracking.first_seen:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tracking_proto_init() }
func file_tracking_proto_init() {
	if File_tracking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AkitaWitnessTracking); i {
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
			RawDescriptor: file_tracking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tracking_proto_goTypes,
		DependencyIndexes: file_tracking_proto_depIdxs,
		MessageInfos:      file_tracking_proto_msgTypes,
	}.Build()
	File_tracking_proto = out.File
	file_tracking_proto_rawDesc = nil
	file_tracking_proto_goTypes = nil
	file_tracking_proto_depIdxs = nil
}
