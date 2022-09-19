// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/terraform/terraform.proto

package api

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListTerraformObjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string      `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Namespace   string      `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Pagination  *Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListTerraformObjectsRequest) Reset() {
	*x = ListTerraformObjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_terraform_terraform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTerraformObjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTerraformObjectsRequest) ProtoMessage() {}

func (x *ListTerraformObjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_terraform_terraform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTerraformObjectsRequest.ProtoReflect.Descriptor instead.
func (*ListTerraformObjectsRequest) Descriptor() ([]byte, []int) {
	return file_api_terraform_terraform_proto_rawDescGZIP(), []int{0}
}

func (x *ListTerraformObjectsRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ListTerraformObjectsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListTerraformObjectsRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListTerraformObjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*TerraformObject    `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	Errors  []*TerraformListError `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ListTerraformObjectsResponse) Reset() {
	*x = ListTerraformObjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_terraform_terraform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTerraformObjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTerraformObjectsResponse) ProtoMessage() {}

func (x *ListTerraformObjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_terraform_terraform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTerraformObjectsResponse.ProtoReflect.Descriptor instead.
func (*ListTerraformObjectsResponse) Descriptor() ([]byte, []int) {
	return file_api_terraform_terraform_proto_rawDescGZIP(), []int{1}
}

func (x *ListTerraformObjectsResponse) GetObjects() []*TerraformObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *ListTerraformObjectsResponse) GetErrors() []*TerraformListError {
	if x != nil {
		return x.Errors
	}
	return nil
}

type GetTerraformObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *GetTerraformObjectRequest) Reset() {
	*x = GetTerraformObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_terraform_terraform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTerraformObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTerraformObjectRequest) ProtoMessage() {}

func (x *GetTerraformObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_terraform_terraform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTerraformObjectRequest.ProtoReflect.Descriptor instead.
func (*GetTerraformObjectRequest) Descriptor() ([]byte, []int) {
	return file_api_terraform_terraform_proto_rawDescGZIP(), []int{2}
}

func (x *GetTerraformObjectRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *GetTerraformObjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTerraformObjectRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type GetTerraformObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *TerraformObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Yaml   string           `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
}

func (x *GetTerraformObjectResponse) Reset() {
	*x = GetTerraformObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_terraform_terraform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTerraformObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTerraformObjectResponse) ProtoMessage() {}

func (x *GetTerraformObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_terraform_terraform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTerraformObjectResponse.ProtoReflect.Descriptor instead.
func (*GetTerraformObjectResponse) Descriptor() ([]byte, []int) {
	return file_api_terraform_terraform_proto_rawDescGZIP(), []int{3}
}

func (x *GetTerraformObjectResponse) GetObject() *TerraformObject {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *GetTerraformObjectResponse) GetYaml() string {
	if x != nil {
		return x.Yaml
	}
	return ""
}

type SyncTerraformObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *SyncTerraformObjectRequest) Reset() {
	*x = SyncTerraformObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_terraform_terraform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncTerraformObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncTerraformObjectRequest) ProtoMessage() {}

func (x *SyncTerraformObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_terraform_terraform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncTerraformObjectRequest.ProtoReflect.Descriptor instead.
func (*SyncTerraformObjectRequest) Descriptor() ([]byte, []int) {
	return file_api_terraform_terraform_proto_rawDescGZIP(), []int{4}
}

func (x *SyncTerraformObjectRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *SyncTerraformObjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SyncTerraformObjectRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type SyncTerraformObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SyncTerraformObjectResponse) Reset() {
	*x = SyncTerraformObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_terraform_terraform_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncTerraformObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncTerraformObjectResponse) ProtoMessage() {}

func (x *SyncTerraformObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_terraform_terraform_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncTerraformObjectResponse.ProtoReflect.Descriptor instead.
func (*SyncTerraformObjectResponse) Descriptor() ([]byte, []int) {
	return file_api_terraform_terraform_proto_rawDescGZIP(), []int{5}
}

func (x *SyncTerraformObjectResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_terraform_terraform_proto protoreflect.FileDescriptor

var file_api_terraform_terraform_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x91, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x22, 0x6f, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x65, 0x72, 0x72, 0x61,
	0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x67, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x54, 0x65, 0x72, 0x72,
	0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61,
	0x6d, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x22, 0x70,
	0x0a, 0x1a, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x22, 0x37, 0x0a, 0x1b, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72,
	0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xbe, 0x03, 0x0a, 0x09, 0x54, 0x65,
	0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x8c, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x12, 0x29, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x74, 0x65,
	0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12,
	0x15, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x27, 0x2e,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x91, 0x01, 0x0a, 0x13, 0x53, 0x79, 0x6e, 0x63, 0x54,
	0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x28,
	0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x65, 0x72, 0x72,
	0x61, 0x66, 0x6f, 0x72, 0x6d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x3a, 0x01, 0x2a, 0x42, 0xbf, 0x01, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2d, 0x67, 0x69, 0x74, 0x6f, 0x70,
	0x73, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x92, 0x41, 0x7f, 0x12, 0x59, 0x0a,
	0x1a, 0x57, 0x65, 0x61, 0x76, 0x65, 0x20, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x20, 0x54, 0x65,
	0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x41, 0x50, 0x49, 0x12, 0x36, 0x54, 0x68, 0x65,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x54,
	0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_terraform_terraform_proto_rawDescOnce sync.Once
	file_api_terraform_terraform_proto_rawDescData = file_api_terraform_terraform_proto_rawDesc
)

func file_api_terraform_terraform_proto_rawDescGZIP() []byte {
	file_api_terraform_terraform_proto_rawDescOnce.Do(func() {
		file_api_terraform_terraform_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_terraform_terraform_proto_rawDescData)
	})
	return file_api_terraform_terraform_proto_rawDescData
}

var file_api_terraform_terraform_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_terraform_terraform_proto_goTypes = []interface{}{
	(*ListTerraformObjectsRequest)(nil),  // 0: terraform.v1.ListTerraformObjectsRequest
	(*ListTerraformObjectsResponse)(nil), // 1: terraform.v1.ListTerraformObjectsResponse
	(*GetTerraformObjectRequest)(nil),    // 2: terraform.v1.GetTerraformObjectRequest
	(*GetTerraformObjectResponse)(nil),   // 3: terraform.v1.GetTerraformObjectResponse
	(*SyncTerraformObjectRequest)(nil),   // 4: terraform.v1.SyncTerraformObjectRequest
	(*SyncTerraformObjectResponse)(nil),  // 5: terraform.v1.SyncTerraformObjectResponse
	(*Pagination)(nil),                   // 6: terraform.v1.Pagination
	(*TerraformObject)(nil),              // 7: terraform.v1.TerraformObject
	(*TerraformListError)(nil),           // 8: terraform.v1.TerraformListError
}
var file_api_terraform_terraform_proto_depIdxs = []int32{
	6, // 0: terraform.v1.ListTerraformObjectsRequest.pagination:type_name -> terraform.v1.Pagination
	7, // 1: terraform.v1.ListTerraformObjectsResponse.objects:type_name -> terraform.v1.TerraformObject
	8, // 2: terraform.v1.ListTerraformObjectsResponse.errors:type_name -> terraform.v1.TerraformListError
	7, // 3: terraform.v1.GetTerraformObjectResponse.object:type_name -> terraform.v1.TerraformObject
	0, // 4: terraform.v1.Terraform.ListTerraformObjects:input_type -> terraform.v1.ListTerraformObjectsRequest
	2, // 5: terraform.v1.Terraform.GetTerraformObject:input_type -> terraform.v1.GetTerraformObjectRequest
	4, // 6: terraform.v1.Terraform.SyncTerraformObject:input_type -> terraform.v1.SyncTerraformObjectRequest
	1, // 7: terraform.v1.Terraform.ListTerraformObjects:output_type -> terraform.v1.ListTerraformObjectsResponse
	3, // 8: terraform.v1.Terraform.GetTerraformObject:output_type -> terraform.v1.GetTerraformObjectResponse
	5, // 9: terraform.v1.Terraform.SyncTerraformObject:output_type -> terraform.v1.SyncTerraformObjectResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_terraform_terraform_proto_init() }
func file_api_terraform_terraform_proto_init() {
	if File_api_terraform_terraform_proto != nil {
		return
	}
	file_api_terraform_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_terraform_terraform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTerraformObjectsRequest); i {
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
		file_api_terraform_terraform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTerraformObjectsResponse); i {
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
		file_api_terraform_terraform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTerraformObjectRequest); i {
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
		file_api_terraform_terraform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTerraformObjectResponse); i {
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
		file_api_terraform_terraform_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncTerraformObjectRequest); i {
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
		file_api_terraform_terraform_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncTerraformObjectResponse); i {
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
			RawDescriptor: file_api_terraform_terraform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_terraform_terraform_proto_goTypes,
		DependencyIndexes: file_api_terraform_terraform_proto_depIdxs,
		MessageInfos:      file_api_terraform_terraform_proto_msgTypes,
	}.Build()
	File_api_terraform_terraform_proto = out.File
	file_api_terraform_terraform_proto_rawDesc = nil
	file_api_terraform_terraform_proto_goTypes = nil
	file_api_terraform_terraform_proto_depIdxs = nil
}
