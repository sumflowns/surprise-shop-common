// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brand.proto

package brand

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//sku的库存
type Brand struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FirstLetter          string   `protobuf:"bytes,3,opt,name=first_letter,json=firstLetter,proto3" json:"first_letter,omitempty"`
	Sort                 int64    `protobuf:"varint,4,opt,name=sort,proto3" json:"sort,omitempty"`
	FactoryStatus        int64    `protobuf:"varint,5,opt,name=factory_status,json=factoryStatus,proto3" json:"factory_status,omitempty"`
	ShowStatus           int64    `protobuf:"varint,6,opt,name=show_status,json=showStatus,proto3" json:"show_status,omitempty"`
	ProductCount         int64    `protobuf:"varint,7,opt,name=product_count,json=productCount,proto3" json:"product_count,omitempty"`
	ProductCommentCount  int64    `protobuf:"varint,8,opt,name=product_comment_count,json=productCommentCount,proto3" json:"product_comment_count,omitempty"`
	Logo                 string   `protobuf:"bytes,9,opt,name=logo,proto3" json:"logo,omitempty"`
	BigPic               string   `protobuf:"bytes,10,opt,name=big_pic,json=bigPic,proto3" json:"big_pic,omitempty"`
	BrandStory           string   `protobuf:"bytes,11,opt,name=brand_story,json=brandStory,proto3" json:"brand_story,omitempty"`
	CreatedTime          string   `protobuf:"bytes,12,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Brand) Reset()         { *m = Brand{} }
func (m *Brand) String() string { return proto.CompactTextString(m) }
func (*Brand) ProtoMessage()    {}
func (*Brand) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{0}
}

func (m *Brand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Brand.Unmarshal(m, b)
}
func (m *Brand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Brand.Marshal(b, m, deterministic)
}
func (m *Brand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Brand.Merge(m, src)
}
func (m *Brand) XXX_Size() int {
	return xxx_messageInfo_Brand.Size(m)
}
func (m *Brand) XXX_DiscardUnknown() {
	xxx_messageInfo_Brand.DiscardUnknown(m)
}

var xxx_messageInfo_Brand proto.InternalMessageInfo

func (m *Brand) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Brand) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Brand) GetFirstLetter() string {
	if m != nil {
		return m.FirstLetter
	}
	return ""
}

func (m *Brand) GetSort() int64 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *Brand) GetFactoryStatus() int64 {
	if m != nil {
		return m.FactoryStatus
	}
	return 0
}

func (m *Brand) GetShowStatus() int64 {
	if m != nil {
		return m.ShowStatus
	}
	return 0
}

func (m *Brand) GetProductCount() int64 {
	if m != nil {
		return m.ProductCount
	}
	return 0
}

func (m *Brand) GetProductCommentCount() int64 {
	if m != nil {
		return m.ProductCommentCount
	}
	return 0
}

func (m *Brand) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *Brand) GetBigPic() string {
	if m != nil {
		return m.BigPic
	}
	return ""
}

func (m *Brand) GetBrandStory() string {
	if m != nil {
		return m.BrandStory
	}
	return ""
}

func (m *Brand) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type In_GetBrandById struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_GetBrandById) Reset()         { *m = In_GetBrandById{} }
func (m *In_GetBrandById) String() string { return proto.CompactTextString(m) }
func (*In_GetBrandById) ProtoMessage()    {}
func (*In_GetBrandById) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{2}
}

func (m *In_GetBrandById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_GetBrandById.Unmarshal(m, b)
}
func (m *In_GetBrandById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_GetBrandById.Marshal(b, m, deterministic)
}
func (m *In_GetBrandById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_GetBrandById.Merge(m, src)
}
func (m *In_GetBrandById) XXX_Size() int {
	return xxx_messageInfo_In_GetBrandById.Size(m)
}
func (m *In_GetBrandById) XXX_DiscardUnknown() {
	xxx_messageInfo_In_GetBrandById.DiscardUnknown(m)
}

var xxx_messageInfo_In_GetBrandById proto.InternalMessageInfo

func (m *In_GetBrandById) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Out_GetBrandById struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Brand                *Brand   `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out_GetBrandById) Reset()         { *m = Out_GetBrandById{} }
func (m *Out_GetBrandById) String() string { return proto.CompactTextString(m) }
func (*Out_GetBrandById) ProtoMessage()    {}
func (*Out_GetBrandById) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{3}
}

func (m *Out_GetBrandById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_GetBrandById.Unmarshal(m, b)
}
func (m *Out_GetBrandById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_GetBrandById.Marshal(b, m, deterministic)
}
func (m *Out_GetBrandById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_GetBrandById.Merge(m, src)
}
func (m *Out_GetBrandById) XXX_Size() int {
	return xxx_messageInfo_Out_GetBrandById.Size(m)
}
func (m *Out_GetBrandById) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_GetBrandById.DiscardUnknown(m)
}

var xxx_messageInfo_Out_GetBrandById proto.InternalMessageInfo

func (m *Out_GetBrandById) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_GetBrandById) GetBrand() *Brand {
	if m != nil {
		return m.Brand
	}
	return nil
}

type Out_UpdateBrandInfo struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Brand                *Brand   `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out_UpdateBrandInfo) Reset()         { *m = Out_UpdateBrandInfo{} }
func (m *Out_UpdateBrandInfo) String() string { return proto.CompactTextString(m) }
func (*Out_UpdateBrandInfo) ProtoMessage()    {}
func (*Out_UpdateBrandInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{4}
}

func (m *Out_UpdateBrandInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_UpdateBrandInfo.Unmarshal(m, b)
}
func (m *Out_UpdateBrandInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_UpdateBrandInfo.Marshal(b, m, deterministic)
}
func (m *Out_UpdateBrandInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_UpdateBrandInfo.Merge(m, src)
}
func (m *Out_UpdateBrandInfo) XXX_Size() int {
	return xxx_messageInfo_Out_UpdateBrandInfo.Size(m)
}
func (m *Out_UpdateBrandInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_UpdateBrandInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Out_UpdateBrandInfo proto.InternalMessageInfo

func (m *Out_UpdateBrandInfo) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_UpdateBrandInfo) GetBrand() *Brand {
	if m != nil {
		return m.Brand
	}
	return nil
}

type In_UpdateBrandInfo struct {
	Brand                *Brand   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_UpdateBrandInfo) Reset()         { *m = In_UpdateBrandInfo{} }
func (m *In_UpdateBrandInfo) String() string { return proto.CompactTextString(m) }
func (*In_UpdateBrandInfo) ProtoMessage()    {}
func (*In_UpdateBrandInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{5}
}

func (m *In_UpdateBrandInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_UpdateBrandInfo.Unmarshal(m, b)
}
func (m *In_UpdateBrandInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_UpdateBrandInfo.Marshal(b, m, deterministic)
}
func (m *In_UpdateBrandInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_UpdateBrandInfo.Merge(m, src)
}
func (m *In_UpdateBrandInfo) XXX_Size() int {
	return xxx_messageInfo_In_UpdateBrandInfo.Size(m)
}
func (m *In_UpdateBrandInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_In_UpdateBrandInfo.DiscardUnknown(m)
}

var xxx_messageInfo_In_UpdateBrandInfo proto.InternalMessageInfo

func (m *In_UpdateBrandInfo) GetBrand() *Brand {
	if m != nil {
		return m.Brand
	}
	return nil
}

type In_GetBrands struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Pages                int64    `protobuf:"varint,2,opt,name=pages,proto3" json:"pages,omitempty"`
	SearchKey            string   `protobuf:"bytes,3,opt,name=search_key,json=searchKey,proto3" json:"search_key,omitempty"`
	StartTime            string   `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string   `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_GetBrands) Reset()         { *m = In_GetBrands{} }
func (m *In_GetBrands) String() string { return proto.CompactTextString(m) }
func (*In_GetBrands) ProtoMessage()    {}
func (*In_GetBrands) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{6}
}

func (m *In_GetBrands) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_GetBrands.Unmarshal(m, b)
}
func (m *In_GetBrands) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_GetBrands.Marshal(b, m, deterministic)
}
func (m *In_GetBrands) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_GetBrands.Merge(m, src)
}
func (m *In_GetBrands) XXX_Size() int {
	return xxx_messageInfo_In_GetBrands.Size(m)
}
func (m *In_GetBrands) XXX_DiscardUnknown() {
	xxx_messageInfo_In_GetBrands.DiscardUnknown(m)
}

var xxx_messageInfo_In_GetBrands proto.InternalMessageInfo

func (m *In_GetBrands) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *In_GetBrands) GetPages() int64 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *In_GetBrands) GetSearchKey() string {
	if m != nil {
		return m.SearchKey
	}
	return ""
}

func (m *In_GetBrands) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *In_GetBrands) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

type Out_GetBrands struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Pages                int64    `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Total                int64    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	BrandList            []*Brand `protobuf:"bytes,5,rep,name=brand_list,json=brandList,proto3" json:"brand_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out_GetBrands) Reset()         { *m = Out_GetBrands{} }
func (m *Out_GetBrands) String() string { return proto.CompactTextString(m) }
func (*Out_GetBrands) ProtoMessage()    {}
func (*Out_GetBrands) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{7}
}

func (m *Out_GetBrands) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_GetBrands.Unmarshal(m, b)
}
func (m *Out_GetBrands) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_GetBrands.Marshal(b, m, deterministic)
}
func (m *Out_GetBrands) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_GetBrands.Merge(m, src)
}
func (m *Out_GetBrands) XXX_Size() int {
	return xxx_messageInfo_Out_GetBrands.Size(m)
}
func (m *Out_GetBrands) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_GetBrands.DiscardUnknown(m)
}

var xxx_messageInfo_Out_GetBrands proto.InternalMessageInfo

func (m *Out_GetBrands) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_GetBrands) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Out_GetBrands) GetPages() int64 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *Out_GetBrands) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Out_GetBrands) GetBrandList() []*Brand {
	if m != nil {
		return m.BrandList
	}
	return nil
}

type In_DeleteBrands struct {
	BrandList            []int64  `protobuf:"varint,1,rep,packed,name=brand_list,json=brandList,proto3" json:"brand_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_DeleteBrands) Reset()         { *m = In_DeleteBrands{} }
func (m *In_DeleteBrands) String() string { return proto.CompactTextString(m) }
func (*In_DeleteBrands) ProtoMessage()    {}
func (*In_DeleteBrands) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{8}
}

func (m *In_DeleteBrands) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_DeleteBrands.Unmarshal(m, b)
}
func (m *In_DeleteBrands) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_DeleteBrands.Marshal(b, m, deterministic)
}
func (m *In_DeleteBrands) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_DeleteBrands.Merge(m, src)
}
func (m *In_DeleteBrands) XXX_Size() int {
	return xxx_messageInfo_In_DeleteBrands.Size(m)
}
func (m *In_DeleteBrands) XXX_DiscardUnknown() {
	xxx_messageInfo_In_DeleteBrands.DiscardUnknown(m)
}

var xxx_messageInfo_In_DeleteBrands proto.InternalMessageInfo

func (m *In_DeleteBrands) GetBrandList() []int64 {
	if m != nil {
		return m.BrandList
	}
	return nil
}

type Out_DeleteBrands struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out_DeleteBrands) Reset()         { *m = Out_DeleteBrands{} }
func (m *Out_DeleteBrands) String() string { return proto.CompactTextString(m) }
func (*Out_DeleteBrands) ProtoMessage()    {}
func (*Out_DeleteBrands) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{9}
}

func (m *Out_DeleteBrands) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_DeleteBrands.Unmarshal(m, b)
}
func (m *Out_DeleteBrands) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_DeleteBrands.Marshal(b, m, deterministic)
}
func (m *Out_DeleteBrands) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_DeleteBrands.Merge(m, src)
}
func (m *Out_DeleteBrands) XXX_Size() int {
	return xxx_messageInfo_Out_DeleteBrands.Size(m)
}
func (m *Out_DeleteBrands) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_DeleteBrands.DiscardUnknown(m)
}

var xxx_messageInfo_Out_DeleteBrands proto.InternalMessageInfo

func (m *Out_DeleteBrands) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type In_CreateBrand struct {
	Brand                *Brand   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_CreateBrand) Reset()         { *m = In_CreateBrand{} }
func (m *In_CreateBrand) String() string { return proto.CompactTextString(m) }
func (*In_CreateBrand) ProtoMessage()    {}
func (*In_CreateBrand) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{10}
}

func (m *In_CreateBrand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_CreateBrand.Unmarshal(m, b)
}
func (m *In_CreateBrand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_CreateBrand.Marshal(b, m, deterministic)
}
func (m *In_CreateBrand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_CreateBrand.Merge(m, src)
}
func (m *In_CreateBrand) XXX_Size() int {
	return xxx_messageInfo_In_CreateBrand.Size(m)
}
func (m *In_CreateBrand) XXX_DiscardUnknown() {
	xxx_messageInfo_In_CreateBrand.DiscardUnknown(m)
}

var xxx_messageInfo_In_CreateBrand proto.InternalMessageInfo

func (m *In_CreateBrand) GetBrand() *Brand {
	if m != nil {
		return m.Brand
	}
	return nil
}

type Out_CreateBrand struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Brand                *Brand   `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out_CreateBrand) Reset()         { *m = Out_CreateBrand{} }
func (m *Out_CreateBrand) String() string { return proto.CompactTextString(m) }
func (*Out_CreateBrand) ProtoMessage()    {}
func (*Out_CreateBrand) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b213c818d79a472, []int{11}
}

func (m *Out_CreateBrand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_CreateBrand.Unmarshal(m, b)
}
func (m *Out_CreateBrand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_CreateBrand.Marshal(b, m, deterministic)
}
func (m *Out_CreateBrand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_CreateBrand.Merge(m, src)
}
func (m *Out_CreateBrand) XXX_Size() int {
	return xxx_messageInfo_Out_CreateBrand.Size(m)
}
func (m *Out_CreateBrand) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_CreateBrand.DiscardUnknown(m)
}

var xxx_messageInfo_Out_CreateBrand proto.InternalMessageInfo

func (m *Out_CreateBrand) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_CreateBrand) GetBrand() *Brand {
	if m != nil {
		return m.Brand
	}
	return nil
}

func init() {
	proto.RegisterType((*Brand)(nil), "brand.Brand")
	proto.RegisterType((*Error)(nil), "brand.Error")
	proto.RegisterType((*In_GetBrandById)(nil), "brand.In_GetBrandById")
	proto.RegisterType((*Out_GetBrandById)(nil), "brand.Out_GetBrandById")
	proto.RegisterType((*Out_UpdateBrandInfo)(nil), "brand.Out_UpdateBrandInfo")
	proto.RegisterType((*In_UpdateBrandInfo)(nil), "brand.In_UpdateBrandInfo")
	proto.RegisterType((*In_GetBrands)(nil), "brand.In_GetBrands")
	proto.RegisterType((*Out_GetBrands)(nil), "brand.Out_GetBrands")
	proto.RegisterType((*In_DeleteBrands)(nil), "brand.In_DeleteBrands")
	proto.RegisterType((*Out_DeleteBrands)(nil), "brand.Out_DeleteBrands")
	proto.RegisterType((*In_CreateBrand)(nil), "brand.In_CreateBrand")
	proto.RegisterType((*Out_CreateBrand)(nil), "brand.Out_CreateBrand")
}

func init() { proto.RegisterFile("brand.proto", fileDescriptor_0b213c818d79a472) }

var fileDescriptor_0b213c818d79a472 = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xce, 0x66, 0xb3, 0x6d, 0x73, 0xb2, 0x6d, 0x65, 0xfa, 0xb7, 0x0d, 0x88, 0xe9, 0x8a, 0x50,
	0x10, 0x8a, 0xd4, 0x1f, 0x7a, 0x25, 0xd8, 0x2a, 0x1a, 0x2d, 0x28, 0xa9, 0x5e, 0x28, 0xc8, 0xb2,
	0xd9, 0x9d, 0xa6, 0x83, 0xbb, 0x3b, 0x61, 0x66, 0x82, 0xe4, 0x39, 0x7c, 0x03, 0x1f, 0xcb, 0xe7,
	0xf0, 0x01, 0x64, 0xce, 0x4c, 0x9a, 0xc9, 0xb6, 0x42, 0x2e, 0x7a, 0x37, 0xe7, 0xfb, 0xbe, 0x73,
	0xf6, 0xcc, 0x77, 0xce, 0x24, 0xd0, 0x19, 0x8a, 0xb4, 0xca, 0x8f, 0xc6, 0x82, 0x2b, 0x4e, 0x02,
	0x0c, 0xe2, 0xbf, 0x4d, 0x08, 0x4e, 0xf5, 0x89, 0x6c, 0x40, 0x93, 0xe5, 0x91, 0xd7, 0xf3, 0x0e,
	0xfd, 0x41, 0x93, 0xe5, 0x84, 0x40, 0xab, 0x4a, 0x4b, 0x1a, 0x35, 0x7b, 0xde, 0x61, 0x7b, 0x80,
	0x67, 0x72, 0x00, 0xe1, 0x25, 0x13, 0x52, 0x25, 0x05, 0x55, 0x8a, 0x8a, 0xc8, 0x47, 0xae, 0x83,
	0xd8, 0x39, 0x42, 0x3a, 0x4d, 0x72, 0xa1, 0xa2, 0x16, 0x16, 0xc2, 0x33, 0x79, 0x04, 0x1b, 0x97,
	0x69, 0xa6, 0xb8, 0x98, 0x26, 0x52, 0xa5, 0x6a, 0x22, 0xa3, 0x00, 0xd9, 0x75, 0x8b, 0x5e, 0x20,
	0x48, 0x1e, 0x40, 0x47, 0x5e, 0xf1, 0x9f, 0x33, 0xcd, 0x0a, 0x6a, 0x40, 0x43, 0x56, 0xf0, 0x10,
	0xd6, 0xc7, 0x82, 0xe7, 0x93, 0x4c, 0x25, 0x19, 0x9f, 0x54, 0x2a, 0x5a, 0x45, 0x49, 0x68, 0xc1,
	0x33, 0x8d, 0x91, 0x63, 0xd8, 0x99, 0x8b, 0xca, 0x92, 0x56, 0x33, 0xf1, 0x1a, 0x8a, 0xb7, 0xae,
	0xc5, 0xc8, 0x99, 0x1c, 0x02, 0xad, 0x82, 0x8f, 0x78, 0xd4, 0x36, 0x77, 0xd5, 0x67, 0xb2, 0x07,
	0xab, 0x43, 0x36, 0x4a, 0xc6, 0x2c, 0x8b, 0x00, 0xe1, 0x95, 0x21, 0x1b, 0x7d, 0x62, 0x99, 0x6e,
	0x13, 0xbd, 0x4b, 0xa4, 0xee, 0x3d, 0xea, 0x20, 0x09, 0x08, 0x5d, 0x68, 0x44, 0xbb, 0x94, 0x09,
	0x9a, 0x2a, 0x9a, 0x27, 0x8a, 0x95, 0x34, 0x0a, 0x8d, 0x4b, 0x16, 0xfb, 0xcc, 0x4a, 0x1a, 0x3f,
	0x87, 0xe0, 0x8d, 0x10, 0x1c, 0xed, 0xca, 0x78, 0x4e, 0xd1, 0xf7, 0x60, 0x80, 0x67, 0x12, 0xc1,
	0x6a, 0x49, 0xa5, 0x4c, 0x47, 0x33, 0xf3, 0x67, 0x61, 0x7c, 0x00, 0x9b, 0xfd, 0x2a, 0x79, 0x4b,
	0x15, 0x8e, 0xec, 0x74, 0xda, 0xbf, 0x31, 0xb6, 0xf8, 0x1b, 0xdc, 0xfb, 0x38, 0x51, 0x8b, 0x9a,
	0x18, 0x02, 0xaa, 0xbf, 0x86, 0xb2, 0xce, 0x71, 0x78, 0x64, 0x16, 0x01, 0x3b, 0x18, 0x18, 0x4a,
	0x6b, 0x10, 0xc5, 0x4f, 0xce, 0x35, 0x58, 0x64, 0x60, 0x97, 0xe5, 0x3b, 0x6c, 0xe9, 0xda, 0x5f,
	0xc6, 0x79, 0xaa, 0x28, 0x32, 0xfd, 0xea, 0x92, 0xdf, 0x59, 0xf9, 0x13, 0x20, 0xfd, 0xea, 0xb6,
	0xea, 0x26, 0xd3, 0xfb, 0x7f, 0xe6, 0x2f, 0x0f, 0x42, 0xc7, 0x18, 0x49, 0xb6, 0x21, 0x28, 0x58,
	0xc9, 0x94, 0x35, 0xc6, 0x04, 0x1a, 0x1d, 0xa7, 0x23, 0x2a, 0xb1, 0x09, 0x7f, 0x60, 0x02, 0x72,
	0x1f, 0x40, 0xd2, 0x54, 0x64, 0x57, 0xc9, 0x0f, 0x3a, 0xb5, 0x2b, 0xdd, 0x36, 0xc8, 0x07, 0x3a,
	0x45, 0x5a, 0xa5, 0x42, 0x99, 0x59, 0xb6, 0x2c, 0xad, 0x11, 0x3d, 0x49, 0xb2, 0x0f, 0x6b, 0xb4,
	0xb2, 0x83, 0x0e, 0xcc, 0xb4, 0x68, 0x65, 0x86, 0xfc, 0xdb, 0x83, 0x75, 0x77, 0x16, 0x72, 0x29,
	0xa7, 0xae, 0x5b, 0x6f, 0xde, 0xda, 0xba, 0xef, 0xb6, 0xbe, 0x0d, 0x81, 0xe2, 0x2a, 0x2d, 0xec,
	0x6b, 0x33, 0x01, 0x79, 0x0c, 0x66, 0x1b, 0x93, 0x82, 0x49, 0x15, 0x05, 0x3d, 0xff, 0x86, 0x6d,
	0x6d, 0x0c, 0xce, 0x99, 0x54, 0xf1, 0x13, 0x5c, 0xa9, 0xd7, 0xb4, 0xa0, 0xd6, 0x74, 0x34, 0xc4,
	0xc9, 0xf7, 0x7a, 0xfe, 0xa1, 0xef, 0x66, 0xbc, 0x30, 0x1b, 0xb6, 0x90, 0xb2, 0xc4, 0xc5, 0xe2,
	0x67, 0xb0, 0xd1, 0xaf, 0x92, 0x33, 0x7c, 0x05, 0xe6, 0x27, 0x67, 0x99, 0xd1, 0x7e, 0x85, 0x4d,
	0xfd, 0xb5, 0x5a, 0xda, 0x5d, 0xec, 0xdb, 0xf1, 0x9f, 0x26, 0x84, 0x08, 0xbc, 0x4b, 0xab, 0xbc,
	0xa0, 0x82, 0xbc, 0x82, 0x70, 0xe1, 0xdd, 0xec, 0xda, 0xac, 0xda, 0x9b, 0xeb, 0xee, 0x59, 0xbc,
	0xfe, 0xd0, 0xe2, 0x06, 0x79, 0x0f, 0x9b, 0xf5, 0x05, 0xde, 0x9f, 0x57, 0xa9, 0x51, 0xdd, 0xae,
	0x53, 0xa8, 0xc6, 0xc5, 0x0d, 0x72, 0x02, 0xed, 0xf9, 0xea, 0x6c, 0xdd, 0xec, 0x45, 0x76, 0xb7,
	0x6f, 0x69, 0x44, 0xc6, 0x0d, 0x7d, 0x91, 0x85, 0xf1, 0x38, 0x17, 0x71, 0xf1, 0x85, 0x8b, 0xb8,
	0x44, 0xdc, 0x20, 0x2f, 0xa1, 0xe3, 0x7a, 0xbe, 0x33, 0xaf, 0xe0, 0xc0, 0xdd, 0x5d, 0xa7, 0x80,
	0x83, 0xc7, 0x8d, 0xe1, 0x0a, 0xfe, 0xcd, 0x3c, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xf4,
	0xf0, 0x41, 0x75, 0x06, 0x00, 0x00,
}
