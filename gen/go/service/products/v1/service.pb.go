// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: service/products/v1/service.proto

package pb_prod_products

import (
	v1 "github.com/viktar3w/service-contracts/gen/go/common/filter/v1"
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

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//Description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	//Image ID
	ImageId string `protobuf:"bytes,3,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	//Price
	Price string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	//Currency ID
	CurrencyId uint32 `protobuf:"varint,5,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	//Rating
	Rating uint32 `protobuf:"varint,6,opt,name=rating,proto3" json:"rating,omitempty"`
	//Category ID
	CategoryId uint32 `protobuf:"varint,7,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	//Specification
	Specification string `protobuf:"bytes,8,opt,name=specification,proto3" json:"specification,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_products_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_products_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_service_products_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *CreateProductRequest) GetCurrencyId() uint32 {
	if x != nil {
		return x.CurrencyId
	}
	return 0
}

func (x *CreateProductRequest) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *CreateProductRequest) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CreateProductRequest) GetSpecification() string {
	if x != nil {
		return x.Specification
	}
	return ""
}

type ProductServiceCreateProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *ProductServiceCreateProductResponse) Reset() {
	*x = ProductServiceCreateProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_products_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductServiceCreateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductServiceCreateProductResponse) ProtoMessage() {}

func (x *ProductServiceCreateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_products_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductServiceCreateProductResponse.ProtoReflect.Descriptor instead.
func (*ProductServiceCreateProductResponse) Descriptor() ([]byte, []int) {
	return file_service_products_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ProductServiceCreateProductResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type AllProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination  *v1.Pagination        `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Name        *v1.StringFieldFilter `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description *v1.StringFieldFilter `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       *v1.StringFieldFilter `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Rating      *v1.IntFieldFilter    `protobuf:"bytes,5,opt,name=rating,proto3" json:"rating,omitempty"`
	CategoryId  *v1.IntFieldFilter    `protobuf:"bytes,6,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *AllProductsRequest) Reset() {
	*x = AllProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_products_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllProductsRequest) ProtoMessage() {}

func (x *AllProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_products_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllProductsRequest.ProtoReflect.Descriptor instead.
func (*AllProductsRequest) Descriptor() ([]byte, []int) {
	return file_service_products_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *AllProductsRequest) GetPagination() *v1.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *AllProductsRequest) GetName() *v1.StringFieldFilter {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *AllProductsRequest) GetDescription() *v1.StringFieldFilter {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *AllProductsRequest) GetPrice() *v1.StringFieldFilter {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *AllProductsRequest) GetRating() *v1.IntFieldFilter {
	if x != nil {
		return x.Rating
	}
	return nil
}

func (x *AllProductsRequest) GetCategoryId() *v1.IntFieldFilter {
	if x != nil {
		return x.CategoryId
	}
	return nil
}

type AllProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product []*Product `protobuf:"bytes,1,rep,name=product,proto3" json:"product,omitempty"`
}

func (x *AllProductsResponse) Reset() {
	*x = AllProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_products_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllProductsResponse) ProtoMessage() {}

func (x *AllProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_products_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllProductsResponse.ProtoReflect.Descriptor instead.
func (*AllProductsResponse) Descriptor() ([]byte, []int) {
	return file_service_products_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *AllProductsResponse) GetProduct() []*Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_service_products_v1_service_proto protoreflect.FileDescriptor

var file_service_products_v1_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x23, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x8a, 0x03, 0x0a, 0x12, 0x41, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x13, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0xec, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x62, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x6b, 0x74, 0x61, 0x72, 0x33, 0x77, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_products_v1_service_proto_rawDescOnce sync.Once
	file_service_products_v1_service_proto_rawDescData = file_service_products_v1_service_proto_rawDesc
)

func file_service_products_v1_service_proto_rawDescGZIP() []byte {
	file_service_products_v1_service_proto_rawDescOnce.Do(func() {
		file_service_products_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_products_v1_service_proto_rawDescData)
	})
	return file_service_products_v1_service_proto_rawDescData
}

var file_service_products_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_products_v1_service_proto_goTypes = []interface{}{
	(*CreateProductRequest)(nil),                // 0: service.products.v1.CreateProductRequest
	(*ProductServiceCreateProductResponse)(nil), // 1: service.products.v1.ProductServiceCreateProductResponse
	(*AllProductsRequest)(nil),                  // 2: service.products.v1.AllProductsRequest
	(*AllProductsResponse)(nil),                 // 3: service.products.v1.AllProductsResponse
	(*Product)(nil),                             // 4: service.products.v1.Product
	(*v1.Pagination)(nil),                       // 5: common.filter.v1.Pagination
	(*v1.StringFieldFilter)(nil),                // 6: common.filter.v1.StringFieldFilter
	(*v1.IntFieldFilter)(nil),                   // 7: common.filter.v1.IntFieldFilter
}
var file_service_products_v1_service_proto_depIdxs = []int32{
	4,  // 0: service.products.v1.ProductServiceCreateProductResponse.product:type_name -> service.products.v1.Product
	5,  // 1: service.products.v1.AllProductsRequest.pagination:type_name -> common.filter.v1.Pagination
	6,  // 2: service.products.v1.AllProductsRequest.name:type_name -> common.filter.v1.StringFieldFilter
	6,  // 3: service.products.v1.AllProductsRequest.description:type_name -> common.filter.v1.StringFieldFilter
	6,  // 4: service.products.v1.AllProductsRequest.price:type_name -> common.filter.v1.StringFieldFilter
	7,  // 5: service.products.v1.AllProductsRequest.rating:type_name -> common.filter.v1.IntFieldFilter
	7,  // 6: service.products.v1.AllProductsRequest.category_id:type_name -> common.filter.v1.IntFieldFilter
	4,  // 7: service.products.v1.AllProductsResponse.product:type_name -> service.products.v1.Product
	0,  // 8: service.products.v1.ProductService.CreateProduct:input_type -> service.products.v1.CreateProductRequest
	2,  // 9: service.products.v1.ProductService.AllProducts:input_type -> service.products.v1.AllProductsRequest
	1,  // 10: service.products.v1.ProductService.CreateProduct:output_type -> service.products.v1.ProductServiceCreateProductResponse
	3,  // 11: service.products.v1.ProductService.AllProducts:output_type -> service.products.v1.AllProductsResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_service_products_v1_service_proto_init() }
func file_service_products_v1_service_proto_init() {
	if File_service_products_v1_service_proto != nil {
		return
	}
	file_service_products_v1_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_products_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_service_products_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductServiceCreateProductResponse); i {
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
		file_service_products_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllProductsRequest); i {
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
		file_service_products_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllProductsResponse); i {
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
			RawDescriptor: file_service_products_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_products_v1_service_proto_goTypes,
		DependencyIndexes: file_service_products_v1_service_proto_depIdxs,
		MessageInfos:      file_service_products_v1_service_proto_msgTypes,
	}.Build()
	File_service_products_v1_service_proto = out.File
	file_service_products_v1_service_proto_rawDesc = nil
	file_service_products_v1_service_proto_goTypes = nil
	file_service_products_v1_service_proto_depIdxs = nil
}
