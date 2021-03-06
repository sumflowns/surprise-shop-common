// Code generated by protoc-gen-go. DO NOT EDIT.
// source: invent.proto

package com_surprise_shop_srv_invent

import (
	_ "brand"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "inventory"
	math "math"
	_ "product"
	_ "product_attribute"
	_ "product_attribute_category"
	_ "product_category"
	_ "product_sku_stock"
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

func init() { proto.RegisterFile("invent.proto", fileDescriptor_5e05fc71e97d0c74) }

var fileDescriptor_5e05fc71e97d0c74 = []byte{
	// 985 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0xcd, 0x6f, 0x1c, 0x35,
	0x14, 0xf7, 0x02, 0x2d, 0xe2, 0x35, 0xb4, 0xf4, 0xa5, 0xd9, 0x25, 0xa6, 0x1f, 0x92, 0x25, 0x3e,
	0xd2, 0xa6, 0xb3, 0xed, 0x96, 0x03, 0x52, 0x21, 0x28, 0x49, 0xa5, 0x74, 0x01, 0x09, 0x44, 0xe0,
	0x1c, 0xed, 0x87, 0x0b, 0xab, 0x6c, 0x67, 0xb6, 0x63, 0x4f, 0xab, 0x16, 0x09, 0x21, 0x71, 0xe3,
	0xc0, 0x3f, 0xc0, 0x09, 0x89, 0x2b, 0x7f, 0x0f, 0x12, 0x37, 0xfe, 0x08, 0xce, 0xa0, 0xf1, 0xd8,
	0x3b, 0xf6, 0xd8, 0xde, 0xd9, 0xc0, 0x25, 0x1b, 0xbd, 0xdf, 0xef, 0x7d, 0xbf, 0x37, 0xb6, 0x61,
	0x63, 0x96, 0x3e, 0xe5, 0xa9, 0x4c, 0x16, 0x79, 0x26, 0x33, 0xbc, 0x3a, 0xc9, 0x1e, 0x27, 0xa2,
	0xc8, 0x17, 0xf9, 0x4c, 0xf0, 0x44, 0x7c, 0x9b, 0x2d, 0x12, 0x91, 0x3f, 0x4d, 0x2a, 0x0e, 0xdd,
	0xae, 0x7e, 0xb3, 0xfc, 0x79, 0x7f, 0xf9, 0x5f, 0xa5, 0x48, 0xb7, 0x16, 0x79, 0x36, 0x2d, 0x26,
	0xb2, 0xaf, 0x7f, 0xb5, 0xf8, 0xf2, 0x38, 0x1f, 0xa5, 0xd3, 0xbe, 0xfa, 0xab, 0x45, 0xef, 0x6a,
	0xc6, 0xc9, 0x64, 0x24, 0xf9, 0x37, 0xa5, 0xad, 0xa6, 0x40, 0x13, 0xef, 0x1b, 0xf9, 0x48, 0xca,
	0x7c, 0x36, 0x2e, 0x24, 0xf7, 0x55, 0x7c, 0x48, 0x2b, 0xef, 0x78, 0x0c, 0x5f, 0xa7, 0x49, 0x15,
	0xa7, 0xc5, 0x89, 0x90, 0xd9, 0xe4, 0xb4, 0xef, 0x49, 0x2a, 0xea, 0xe0, 0xe7, 0x97, 0xe1, 0xe2,
	0x17, 0x15, 0xf6, 0x70, 0x94, 0x4e, 0xe7, 0x3c, 0xc7, 0x4f, 0xe1, 0xe2, 0x11, 0x97, 0x5a, 0x78,
	0xf0, 0x7c, 0x38, 0x45, 0x9a, 0x98, 0x1a, 0x0c, 0xd3, 0x13, 0x17, 0xa3, 0x6f, 0x2d, 0xb1, 0xcf,
	0x0b, 0xd9, 0x00, 0x19, 0xc1, 0xaf, 0xe0, 0xf2, 0xd7, 0x8b, 0xe9, 0x48, 0x72, 0x2d, 0x1e, 0xa6,
	0x8f, 0x32, 0xbc, 0x66, 0xdb, 0xf3, 0x60, 0x7a, 0xdd, 0x31, 0xe9, 0xe1, 0x8c, 0xe0, 0x01, 0x5c,
	0xa8, 0x3d, 0x09, 0xec, 0x85, 0xe3, 0x13, 0xf4, 0xcd, 0x48, 0x70, 0x82, 0x91, 0x32, 0xcd, 0x07,
	0x7c, 0xce, 0x97, 0xa6, 0x85, 0x9b, 0xa6, 0x8b, 0x35, 0xd2, 0x74, 0x41, 0x46, 0xf0, 0x21, 0xbc,
	0x7e, 0x98, 0xf3, 0x3a, 0x4e, 0xdc, 0xb6, 0x6d, 0x39, 0x10, 0xa5, 0x8e, 0x29, 0x07, 0x63, 0x64,
	0xf0, 0x1d, 0xbc, 0x31, 0x34, 0x93, 0x68, 0x3a, 0x72, 0x17, 0x5e, 0x39, 0xe6, 0xf3, 0x39, 0x62,
	0x52, 0x0f, 0xe9, 0x97, 0xfc, 0x49, 0xc1, 0x85, 0xa4, 0x9b, 0x8e, 0x4c, 0x2c, 0xb2, 0x54, 0x70,
	0x46, 0xf0, 0x7d, 0x78, 0xf5, 0x30, 0x4b, 0x1f, 0xcd, 0xf2, 0xc7, 0x67, 0xd0, 0x1a, 0xfc, 0xf9,
	0x12, 0x6c, 0x1c, 0x94, 0x93, 0x6d, 0x3c, 0xef, 0xc3, 0xc6, 0x11, 0x97, 0x4a, 0xa4, 0x26, 0xa1,
	0x9b, 0x54, 0x83, 0x5f, 0xd5, 0x79, 0x29, 0xa7, 0x3d, 0x2d, 0xd7, 0x65, 0x5e, 0x02, 0x8c, 0xe0,
	0x27, 0x70, 0xa9, 0x6a, 0xa1, 0x12, 0xaa, 0xfe, 0x6f, 0xd7, 0x56, 0x1a, 0x10, 0xa5, 0x96, 0xa1,
	0x06, 0xc6, 0x08, 0x7e, 0x00, 0xaf, 0x19, 0xeb, 0x02, 0x37, 0xfd, 0x58, 0x04, 0xbd, 0x12, 0x08,
	0xa4, 0x6c, 0xd0, 0x3e, 0x6c, 0x54, 0x4d, 0xd3, 0xca, 0x56, 0x22, 0xb6, 0xdc, 0x49, 0xc4, 0x06,
	0x18, 0xc1, 0x3d, 0xb8, 0x50, 0x35, 0x4b, 0x49, 0x70, 0xab, 0xb6, 0x60, 0x89, 0x69, 0xd7, 0x32,
	0x60, 0xc9, 0x19, 0x19, 0xfc, 0x7d, 0x0e, 0xba, 0xba, 0xcf, 0x87, 0x7a, 0xb5, 0x4d, 0x99, 0x0b,
	0xe8, 0xd6, 0xc3, 0x69, 0x40, 0x55, 0xf0, 0x5b, 0x89, 0xf7, 0x2d, 0x71, 0x66, 0xdc, 0x26, 0xd3,
	0x5d, 0x9f, 0xec, 0xce, 0xbd, 0xcd, 0x66, 0x04, 0xbf, 0x87, 0x6d, 0x67, 0xbb, 0x0c, 0xac, 0x9a,
	0x94, 0x04, 0x3d, 0x47, 0xf9, 0xb4, 0x1f, 0x76, 0x1e, 0x55, 0x60, 0x04, 0xe7, 0xb0, 0xe9, 0xc7,
	0x26, 0xf0, 0xbd, 0x35, 0x73, 0x16, 0x74, 0x67, 0xdd, 0x84, 0xcb, 0xfe, 0x15, 0xd0, 0x75, 0xf6,
	0xb6, 0x76, 0x18, 0x2e, 0x72, 0x98, 0x1c, 0x2b, 0x72, 0x98, 0xcd, 0x08, 0xe6, 0xb0, 0xe5, 0xec,
	0xb8, 0xc1, 0xf0, 0x66, 0xd0, 0x6b, 0x90, 0x4b, 0x6f, 0x85, 0x9d, 0x06, 0xc9, 0x8c, 0xe0, 0x4f,
	0x1d, 0xb8, 0xe6, 0x17, 0xe1, 0x01, 0x97, 0xa3, 0xd9, 0x5c, 0x7c, 0x36, 0x13, 0x12, 0x07, 0x6b,
	0xd6, 0xd8, 0xd2, 0xa1, 0xf7, 0xd6, 0xad, 0xb6, 0xa5, 0xc4, 0xc8, 0xe0, 0xaf, 0xf3, 0x70, 0x43,
	0x13, 0xf6, 0xcd, 0x41, 0xd5, 0x5c, 0x80, 0xdf, 0x3a, 0x70, 0xa3, 0xb6, 0xe3, 0xd1, 0xd4, 0x2a,
	0xec, 0x25, 0x2b, 0xce, 0x48, 0x27, 0xf8, 0xa0, 0x3e, 0xfd, 0x78, 0x95, 0xbe, 0x9b, 0x48, 0xd0,
	0x00, 0x23, 0xf8, 0x7b, 0x07, 0x98, 0x33, 0xd0, 0x1e, 0x51, 0xad, 0xce, 0x7e, 0x4b, 0xa4, 0xed,
	0x26, 0xe8, 0x41, 0x5b, 0xb0, 0xed, 0x36, 0x18, 0xc1, 0x5f, 0x3a, 0x70, 0x75, 0x45, 0x56, 0x02,
	0xef, 0xff, 0xf7, 0x9a, 0x0a, 0xfa, 0xe1, 0xff, 0x28, 0x68, 0xb9, 0x19, 0x65, 0xd3, 0x9d, 0xb5,
	0x09, 0x04, 0xd8, 0xd6, 0xf4, 0x16, 0xfd, 0xf6, 0xa6, 0xb7, 0x18, 0x60, 0x04, 0x7f, 0xed, 0xc0,
	0x75, 0x67, 0xd1, 0x3c, 0x16, 0x7e, 0xd4, 0x12, 0xe5, 0x6a, 0x75, 0xba, 0xd7, 0x16, 0xe4, 0x6a,
	0x7d, 0x46, 0x06, 0xff, 0x9c, 0x83, 0x5e, 0x13, 0x36, 0xbb, 0xf5, 0x02, 0x7a, 0x81, 0x46, 0xa8,
	0x95, 0xba, 0xed, 0x3b, 0x8e, 0x74, 0x5d, 0x6d, 0x50, 0x12, 0xa0, 0x47, 0xfa, 0xac, 0x17, 0xe6,
	0xc7, 0x0e, 0xd0, 0xf0, 0xa4, 0xaa, 0x45, 0xb9, 0x13, 0xf6, 0x1f, 0xd7, 0xa0, 0x77, 0x23, 0x21,
	0xc4, 0x55, 0x18, 0xc1, 0x27, 0x70, 0x25, 0x10, 0xa2, 0xb0, 0xbe, 0xc0, 0x6d, 0xe9, 0x0b, 0xeb,
	0x0b, 0xdc, 0x9a, 0x7b, 0x39, 0x34, 0x2f, 0xa0, 0x17, 0x9e, 0x2c, 0x11, 0x2b, 0x7a, 0x84, 0x1e,
	0x2d, 0x7a, 0x84, 0xcf, 0x08, 0x3e, 0x83, 0x6e, 0x78, 0x60, 0x70, 0x37, 0xec, 0x3a, 0xcc, 0xa6,
	0xb7, 0x23, 0x9e, 0xc3, 0xf4, 0x2a, 0xe9, 0x40, 0x39, 0xd4, 0x79, 0xb3, 0xfe, 0xa4, 0xa9, 0xa3,
	0xe6, 0x0c, 0x93, 0xa6, 0x4f, 0x99, 0x3f, 0xea, 0xdb, 0xd5, 0xf1, 0x69, 0x71, 0x5c, 0x3e, 0x71,
	0xcc, 0x02, 0x3c, 0xb3, 0x6f, 0x57, 0x06, 0x54, 0xf3, 0x5f, 0xd7, 0xa3, 0x7e, 0x17, 0x39, 0x51,
	0xd9, 0x6c, 0xab, 0x1e, 0x35, 0xdb, 0x0d, 0xca, 0xa6, 0x33, 0x82, 0x3f, 0x74, 0x1a, 0x17, 0x2c,
	0x83, 0xab, 0xe1, 0xef, 0x87, 0x9d, 0x47, 0x15, 0xe8, 0x9d, 0x88, 0xff, 0xa8, 0x06, 0x23, 0x98,
	0xda, 0x57, 0x2c, 0x83, 0x09, 0xdc, 0x59, 0x37, 0x71, 0x41, 0x6f, 0xae, 0x9d, 0xb5, 0x9e, 0x3d,
	0x67, 0x30, 0x6b, 0x97, 0x91, 0x5a, 0x87, 0xd9, 0xd1, 0x5a, 0x87, 0xe9, 0x8c, 0xa0, 0x6c, 0x5c,
	0xb3, 0x0c, 0x66, 0x5d, 0xee, 0x1c, 0xbf, 0x41, 0x32, 0xdd, 0x8d, 0xb8, 0x0d, 0xb2, 0x19, 0xc1,
	0x31, 0x5c, 0x3a, 0xe2, 0xb2, 0x7c, 0x9c, 0x2d, 0xfd, 0xbd, 0x1d, 0x2d, 0xad, 0x4d, 0xa3, 0xef,
	0xc4, 0xcb, 0x6a, 0xf3, 0x18, 0x19, 0x9f, 0x57, 0x2f, 0xf5, 0x7b, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0x41, 0x7a, 0xf8, 0xd8, 0x10, 0x00, 0x00,
}
