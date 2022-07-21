// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: api/rule/v1/error.proto

package v1

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

// @plugins=protoc-gen-go-errors
// 错误
type Error int32

const (
	// @msg=未知类型
	// @code=UNKNOWN
	Error_ERR_UNKNOWN Error = 0
	// @msg=成功
	// @code=OK
	Error_ERR_OK_STATUS Error = 1
	// @msg=未找到资源
	// @code=NOT_FOUND
	Error_ERR_NOT_FOUND Error = 2
	// @msg=请求参数无效
	// @code=INVALID_ARGUMENT
	Error_ERR_INVALID_ARGUMENT Error = 3
	// @msg=请求后端存储错误
	// @code=INTERNAL
	Error_ERR_INTERNAL_STORE Error = 4
	// @msg=内部错误
	// @code=INTERNAL
	Error_ERR_INTERNAL_ERROR Error = 5
	// @msg=未找到对应规则
	// @code=NOT_FOUND
	Error_ERR_RULE_NOT_FOUND Error = 6
	// @msg=请确保用户对该资源拥有足够的权限
	// @code=PERMISSION_DENIED
	Error_ERR_FORBIDDEN Error = 7
	// @msg=请确保用户权限
	// @code=PERMISSION_DENIED
	Error_ERR_UNAUTHORIZED Error = 8
	// @msg=建立连接失败
	// @code=OK
	Error_ERR_FAILED_KAFKA_CONNECTION Error = 9
	// @msg=成功
	// @code=OK
	Error_ERR_OK_KAFKA_CONNECTION Error = 10
	// @msg=重复创建
	// @code=INVALID_ARGUMENT
	Error_ERR_DUPLICATE_CREATE Error = 11
	// @msg=不能删除正在运行的规则
	// @code=INVALID_ARGUMENT
	Error_ERR_CANT_DELETE_RUNNING_RULE Error = 12
	// @msg=建立连接失败
	// @code=OK
	Error_ERR_FAILED_MYSQL_CONNECTION Error = 13
	// @msg=建立连接失败
	// @code=OK
	Error_ERR_FAILED_CLICKHOUSE_CONNECTION Error = 14
	// @msg=获取配置信息失败
	// @code=OK
	Error_ERR_FAILED_SINK_INFO Error = 15
	// @msg=获取映射信息失败
	// @code=OK
	Error_ERR_FAILED_MAP_INFO Error = 16
	// @msg=获取数据表信息失败
	// @code=OK
	Error_ERR_FAILED_TABLE_INFO Error = 17
	// @msg=命名重复
	// @code=INVALID_ARGUMENT
	Error_ERR_DUPLICATE_NAME Error = 18
	// @msg=重复添加设备
	// @code=INVALID_ARGUMENT
	Error_ERR_DUPLICATE_DEVICE Error = 19
	// @msg=没有可用的转发
	// @code=INVALID_ARGUMENT
	Error_ERR_INVALID_RULE Error = 20
	// @msg=建立连接失败
	// @code=OK
	Error_ERR_FAILED_INFLUXDB_CONNECTION Error = 21
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0:  "ERR_UNKNOWN",
		1:  "ERR_OK_STATUS",
		2:  "ERR_NOT_FOUND",
		3:  "ERR_INVALID_ARGUMENT",
		4:  "ERR_INTERNAL_STORE",
		5:  "ERR_INTERNAL_ERROR",
		6:  "ERR_RULE_NOT_FOUND",
		7:  "ERR_FORBIDDEN",
		8:  "ERR_UNAUTHORIZED",
		9:  "ERR_FAILED_KAFKA_CONNECTION",
		10: "ERR_OK_KAFKA_CONNECTION",
		11: "ERR_DUPLICATE_CREATE",
		12: "ERR_CANT_DELETE_RUNNING_RULE",
		13: "ERR_FAILED_MYSQL_CONNECTION",
		14: "ERR_FAILED_CLICKHOUSE_CONNECTION",
		15: "ERR_FAILED_SINK_INFO",
		16: "ERR_FAILED_MAP_INFO",
		17: "ERR_FAILED_TABLE_INFO",
		18: "ERR_DUPLICATE_NAME",
		19: "ERR_DUPLICATE_DEVICE",
		20: "ERR_INVALID_RULE",
		21: "ERR_FAILED_INFLUXDB_CONNECTION",
	}
	Error_value = map[string]int32{
		"ERR_UNKNOWN":                      0,
		"ERR_OK_STATUS":                    1,
		"ERR_NOT_FOUND":                    2,
		"ERR_INVALID_ARGUMENT":             3,
		"ERR_INTERNAL_STORE":               4,
		"ERR_INTERNAL_ERROR":               5,
		"ERR_RULE_NOT_FOUND":               6,
		"ERR_FORBIDDEN":                    7,
		"ERR_UNAUTHORIZED":                 8,
		"ERR_FAILED_KAFKA_CONNECTION":      9,
		"ERR_OK_KAFKA_CONNECTION":          10,
		"ERR_DUPLICATE_CREATE":             11,
		"ERR_CANT_DELETE_RUNNING_RULE":     12,
		"ERR_FAILED_MYSQL_CONNECTION":      13,
		"ERR_FAILED_CLICKHOUSE_CONNECTION": 14,
		"ERR_FAILED_SINK_INFO":             15,
		"ERR_FAILED_MAP_INFO":              16,
		"ERR_FAILED_TABLE_INFO":            17,
		"ERR_DUPLICATE_NAME":               18,
		"ERR_DUPLICATE_DEVICE":             19,
		"ERR_INVALID_RULE":                 20,
		"ERR_FAILED_INFLUXDB_CONNECTION":   21,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_api_rule_v1_error_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_api_rule_v1_error_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_api_rule_v1_error_proto_rawDescGZIP(), []int{0}
}

var File_api_rule_v1_error_proto protoreflect.FileDescriptor

var file_api_rule_v1_error_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2a, 0xc4, 0x04, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b,
	0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x52, 0x52, 0x5f, 0x4f, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x16, 0x0a,
	0x12, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x54,
	0x4f, 0x52, 0x45, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x16, 0x0a,
	0x12, 0x45, 0x52, 0x52, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x5f, 0x46, 0x4f, 0x52,
	0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x5f,
	0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x08, 0x12, 0x1f,
	0x0a, 0x1b, 0x45, 0x52, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x4b, 0x41, 0x46,
	0x4b, 0x41, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12,
	0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x4f, 0x4b, 0x5f, 0x4b, 0x41, 0x46, 0x4b, 0x41, 0x5f,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14,
	0x45, 0x52, 0x52, 0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x10, 0x0b, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x41,
	0x4e, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e,
	0x47, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x10, 0x0c, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x5f, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x52, 0x52,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x48, 0x4f, 0x55,
	0x53, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0e, 0x12,
	0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x53, 0x49,
	0x4e, 0x4b, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x0f, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x52, 0x52,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x4d, 0x41, 0x50, 0x5f, 0x49, 0x4e, 0x46, 0x4f,
	0x10, 0x10, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x11, 0x12, 0x16, 0x0a,
	0x12, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x10, 0x12, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x55, 0x50,
	0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x13, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52,
	0x55, 0x4c, 0x45, 0x10, 0x14, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x52, 0x52, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x46, 0x4c, 0x55, 0x58, 0x44, 0x42, 0x5f, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x15, 0x42, 0x38, 0x0a, 0x07, 0x72, 0x75, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x72, 0x75, 0x6c, 0x65,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rule_v1_error_proto_rawDescOnce sync.Once
	file_api_rule_v1_error_proto_rawDescData = file_api_rule_v1_error_proto_rawDesc
)

func file_api_rule_v1_error_proto_rawDescGZIP() []byte {
	file_api_rule_v1_error_proto_rawDescOnce.Do(func() {
		file_api_rule_v1_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rule_v1_error_proto_rawDescData)
	})
	return file_api_rule_v1_error_proto_rawDescData
}

var file_api_rule_v1_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_rule_v1_error_proto_goTypes = []interface{}{
	(Error)(0), // 0: rule.v1.Error
}
var file_api_rule_v1_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_rule_v1_error_proto_init() }
func file_api_rule_v1_error_proto_init() {
	if File_api_rule_v1_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_rule_v1_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_rule_v1_error_proto_goTypes,
		DependencyIndexes: file_api_rule_v1_error_proto_depIdxs,
		EnumInfos:         file_api_rule_v1_error_proto_enumTypes,
	}.Build()
	File_api_rule_v1_error_proto = out.File
	file_api_rule_v1_error_proto_rawDesc = nil
	file_api_rule_v1_error_proto_goTypes = nil
	file_api_rule_v1_error_proto_depIdxs = nil
}
