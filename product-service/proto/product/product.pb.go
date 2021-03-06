// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product/product.proto

package go_micro_srv_product

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Product struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Available            int32    `protobuf:"varint,3,opt,name=available,proto3" json:"available,omitempty"`
	Quantity             int32    `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	FarmerId             string   `protobuf:"bytes,5,opt,name=farmer_id,json=farmerId,proto3" json:"farmer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f9f5dc4dd6fa6d7, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetAvailable() int32 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *Product) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *Product) GetFarmerId() string {
	if m != nil {
		return m.FarmerId
	}
	return ""
}

type Image struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f9f5dc4dd6fa6d7, []int{1}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Image) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type QueryParams struct {
	FarmerId             string   `protobuf:"bytes,1,opt,name=farmer_id,json=farmerId,proto3" json:"farmer_id,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryParams) Reset()         { *m = QueryParams{} }
func (m *QueryParams) String() string { return proto.CompactTextString(m) }
func (*QueryParams) ProtoMessage()    {}
func (*QueryParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f9f5dc4dd6fa6d7, []int{2}
}

func (m *QueryParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryParams.Unmarshal(m, b)
}
func (m *QueryParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryParams.Marshal(b, m, deterministic)
}
func (m *QueryParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParams.Merge(m, src)
}
func (m *QueryParams) XXX_Size() int {
	return xxx_messageInfo_QueryParams.Size(m)
}
func (m *QueryParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParams.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParams proto.InternalMessageInfo

func (m *QueryParams) GetFarmerId() string {
	if m != nil {
		return m.FarmerId
	}
	return ""
}

func (m *QueryParams) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type Response struct {
	Product              *Product   `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Products             []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f9f5dc4dd6fa6d7, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *Response) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*Product)(nil), "go.micro.srv.product.Product")
	proto.RegisterType((*Image)(nil), "go.micro.srv.product.Image")
	proto.RegisterType((*QueryParams)(nil), "go.micro.srv.product.QueryParams")
	proto.RegisterType((*Response)(nil), "go.micro.srv.product.Response")
}

func init() { proto.RegisterFile("product/product.proto", fileDescriptor_5f9f5dc4dd6fa6d7) }

var fileDescriptor_5f9f5dc4dd6fa6d7 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcf, 0x4b, 0xc3, 0x30,
	0x18, 0x35, 0xdd, 0xaf, 0xf6, 0x1b, 0x0c, 0x0c, 0x0a, 0x61, 0xfe, 0x60, 0xf6, 0xb4, 0x53, 0x84,
	0x79, 0x10, 0x4f, 0xde, 0x36, 0x76, 0x18, 0xcc, 0xee, 0xe2, 0x4d, 0xb2, 0x25, 0xce, 0xc0, 0xda,
	0xd4, 0x34, 0x2b, 0xf4, 0xe2, 0xc9, 0x7f, 0xca, 0xff, 0x4e, 0x96, 0xa6, 0xd3, 0x42, 0xc1, 0x8b,
	0xa7, 0xbc, 0xbc, 0xef, 0xcb, 0xfb, 0x5e, 0x5e, 0x02, 0xe7, 0xa9, 0x56, 0x7c, 0xbf, 0x31, 0xb7,
	0x6e, 0xa5, 0xa9, 0x56, 0x46, 0xe1, 0xb3, 0xad, 0xa2, 0xb1, 0xdc, 0x68, 0x45, 0x33, 0x9d, 0x53,
	0x57, 0x0b, 0x3f, 0x11, 0xf4, 0x96, 0x25, 0xc6, 0x03, 0xf0, 0x24, 0x27, 0x68, 0x84, 0xc6, 0x41,
	0xe4, 0x49, 0x8e, 0x31, 0xb4, 0x13, 0x16, 0x0b, 0xe2, 0x59, 0xc6, 0x62, 0x7c, 0x09, 0x01, 0xcb,
	0x99, 0xdc, 0xb1, 0xf5, 0x4e, 0x90, 0xd6, 0x08, 0x8d, 0x3b, 0xd1, 0x0f, 0x81, 0x87, 0xe0, 0xbf,
	0xef, 0x59, 0x62, 0xa4, 0x29, 0x48, 0xdb, 0x16, 0x8f, 0x7b, 0x7c, 0x01, 0xc1, 0x2b, 0xd3, 0xb1,
	0xd0, 0x2f, 0x92, 0x93, 0x8e, 0x95, 0xf4, 0x4b, 0x62, 0xce, 0xc3, 0x47, 0xe8, 0xcc, 0x63, 0xb6,
	0x15, 0x4d, 0x1e, 0x52, 0x66, 0xde, 0x2a, 0x0f, 0x07, 0x7c, 0xe0, 0x4c, 0x91, 0x96, 0xe3, 0x83,
	0xc8, 0xe2, 0x70, 0x0a, 0xfd, 0xa7, 0xbd, 0xd0, 0xc5, 0x92, 0x69, 0x16, 0x67, 0xf5, 0x61, 0xa8,
	0x3e, 0xac, 0xe6, 0xb2, 0x55, 0x77, 0x19, 0x7e, 0x80, 0x1f, 0x89, 0x2c, 0x55, 0x49, 0x26, 0xf0,
	0x3d, 0xf4, 0x5c, 0x4c, 0x56, 0xa2, 0x3f, 0xb9, 0xa2, 0x4d, 0x19, 0x52, 0x97, 0x5f, 0x54, 0x75,
	0xe3, 0x07, 0xf0, 0x1d, 0xcc, 0x88, 0x37, 0x6a, 0xfd, 0x7d, 0xf2, 0xd8, 0x3e, 0xf9, 0x42, 0x30,
	0x70, 0xec, 0x4a, 0xe8, 0x5c, 0x6e, 0x04, 0x7e, 0x86, 0xd3, 0xa9, 0x4c, 0xf8, 0xd4, 0xda, 0xaf,
	0xde, 0xea, 0xa6, 0x59, 0xf0, 0x57, 0x06, 0xc3, 0xeb, 0xe6, 0x96, 0xea, 0x7a, 0xe1, 0x09, 0x5e,
	0x40, 0x77, 0x26, 0xcc, 0x62, 0x35, 0xfb, 0x17, 0xb9, 0x75, 0xd7, 0x7e, 0xb4, 0xbb, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd8, 0x2f, 0x1f, 0x32, 0x81, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ProductService service

type ProductServiceClient interface {
	FindFarmerProduct(ctx context.Context, in *QueryParams, opts ...client.CallOption) (*Response, error)
	GetMSG(ctx context.Context, in *QueryParams, opts ...client.CallOption) (*Response, error)
}

type productServiceClient struct {
	c           client.Client
	serviceName string
}

func NewProductServiceClient(serviceName string, c client.Client) ProductServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.product"
	}
	return &productServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *productServiceClient) FindFarmerProduct(ctx context.Context, in *QueryParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ProductService.FindFarmerProduct", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetMSG(ctx context.Context, in *QueryParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ProductService.GetMSG", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductService service

type ProductServiceHandler interface {
	FindFarmerProduct(context.Context, *QueryParams, *Response) error
	GetMSG(context.Context, *QueryParams, *Response) error
}

func RegisterProductServiceHandler(s server.Server, hdlr ProductServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ProductService{hdlr}, opts...))
}

type ProductService struct {
	ProductServiceHandler
}

func (h *ProductService) FindFarmerProduct(ctx context.Context, in *QueryParams, out *Response) error {
	return h.ProductServiceHandler.FindFarmerProduct(ctx, in, out)
}

func (h *ProductService) GetMSG(ctx context.Context, in *QueryParams, out *Response) error {
	return h.ProductServiceHandler.GetMSG(ctx, in, out)
}
