// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: stockpb/stock.proto

package stockpb

import (
	orderpb "github.com/leebrouse/Gorder/internal/common/genproto/orderpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetItemsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ItemIDs       []string               `protobuf:"bytes,1,rep,name=ItemIDs,proto3" json:"ItemIDs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetItemsRequest) Reset() {
	*x = GetItemsRequest{}
	mi := &file_stockpb_stock_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsRequest) ProtoMessage() {}

func (x *GetItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stockpb_stock_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsRequest.ProtoReflect.Descriptor instead.
func (*GetItemsRequest) Descriptor() ([]byte, []int) {
	return file_stockpb_stock_proto_rawDescGZIP(), []int{0}
}

func (x *GetItemsRequest) GetItemIDs() []string {
	if x != nil {
		return x.ItemIDs
	}
	return nil
}

type GetItemsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*orderpb.Item        `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetItemsResponse) Reset() {
	*x = GetItemsResponse{}
	mi := &file_stockpb_stock_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsResponse) ProtoMessage() {}

func (x *GetItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stockpb_stock_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsResponse.ProtoReflect.Descriptor instead.
func (*GetItemsResponse) Descriptor() ([]byte, []int) {
	return file_stockpb_stock_proto_rawDescGZIP(), []int{1}
}

func (x *GetItemsResponse) GetItems() []*orderpb.Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type CheckIfItemsInStockRequest struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Items         []*orderpb.ItemWithQuantity `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckIfItemsInStockRequest) Reset() {
	*x = CheckIfItemsInStockRequest{}
	mi := &file_stockpb_stock_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckIfItemsInStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfItemsInStockRequest) ProtoMessage() {}

func (x *CheckIfItemsInStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stockpb_stock_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfItemsInStockRequest.ProtoReflect.Descriptor instead.
func (*CheckIfItemsInStockRequest) Descriptor() ([]byte, []int) {
	return file_stockpb_stock_proto_rawDescGZIP(), []int{2}
}

func (x *CheckIfItemsInStockRequest) GetItems() []*orderpb.ItemWithQuantity {
	if x != nil {
		return x.Items
	}
	return nil
}

type CheckIfItemsInStockResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InStock       int32                  `protobuf:"varint,1,opt,name=InStock,proto3" json:"InStock,omitempty"`
	Items         []*orderpb.Item        `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckIfItemsInStockResponse) Reset() {
	*x = CheckIfItemsInStockResponse{}
	mi := &file_stockpb_stock_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckIfItemsInStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfItemsInStockResponse) ProtoMessage() {}

func (x *CheckIfItemsInStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stockpb_stock_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfItemsInStockResponse.ProtoReflect.Descriptor instead.
func (*CheckIfItemsInStockResponse) Descriptor() ([]byte, []int) {
	return file_stockpb_stock_proto_rawDescGZIP(), []int{3}
}

func (x *CheckIfItemsInStockResponse) GetInStock() int32 {
	if x != nil {
		return x.InStock
	}
	return 0
}

func (x *CheckIfItemsInStockResponse) GetItems() []*orderpb.Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_stockpb_stock_proto protoreflect.FileDescriptor

const file_stockpb_stock_proto_rawDesc = "" +
	"\n" +
	"\x13stockpb/stock.proto\x12\astockpb\x1a\x13orderpb/order.proto\"+\n" +
	"\x0fGetItemsRequest\x12\x18\n" +
	"\aItemIDs\x18\x01 \x03(\tR\aItemIDs\"7\n" +
	"\x10GetItemsResponse\x12#\n" +
	"\x05Items\x18\x01 \x03(\v2\r.orderpb.ItemR\x05Items\"M\n" +
	"\x1aCheckIfItemsInStockRequest\x12/\n" +
	"\x05Items\x18\x01 \x03(\v2\x19.orderpb.ItemWithQuantityR\x05Items\"\\\n" +
	"\x1bCheckIfItemsInStockResponse\x12\x18\n" +
	"\aInStock\x18\x01 \x01(\x05R\aInStock\x12#\n" +
	"\x05Items\x18\x02 \x03(\v2\r.orderpb.ItemR\x05Items2\xb1\x01\n" +
	"\fStockService\x12?\n" +
	"\bGetItems\x12\x18.stockpb.GetItemsRequest\x1a\x19.stockpb.GetItemsResponse\x12`\n" +
	"\x13CheckIfItemsInStock\x12#.stockpb.CheckIfItemsInStockRequest\x1a$.stockpb.CheckIfItemsInStockResponseB>Z<github.com/leebrouse/Gorder/internal/common/genproto/stockpbb\x06proto3"

var (
	file_stockpb_stock_proto_rawDescOnce sync.Once
	file_stockpb_stock_proto_rawDescData []byte
)

func file_stockpb_stock_proto_rawDescGZIP() []byte {
	file_stockpb_stock_proto_rawDescOnce.Do(func() {
		file_stockpb_stock_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_stockpb_stock_proto_rawDesc), len(file_stockpb_stock_proto_rawDesc)))
	})
	return file_stockpb_stock_proto_rawDescData
}

var file_stockpb_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stockpb_stock_proto_goTypes = []any{
	(*GetItemsRequest)(nil),             // 0: stockpb.GetItemsRequest
	(*GetItemsResponse)(nil),            // 1: stockpb.GetItemsResponse
	(*CheckIfItemsInStockRequest)(nil),  // 2: stockpb.CheckIfItemsInStockRequest
	(*CheckIfItemsInStockResponse)(nil), // 3: stockpb.CheckIfItemsInStockResponse
	(*orderpb.Item)(nil),                // 4: orderpb.Item
	(*orderpb.ItemWithQuantity)(nil),    // 5: orderpb.ItemWithQuantity
}
var file_stockpb_stock_proto_depIdxs = []int32{
	4, // 0: stockpb.GetItemsResponse.Items:type_name -> orderpb.Item
	5, // 1: stockpb.CheckIfItemsInStockRequest.Items:type_name -> orderpb.ItemWithQuantity
	4, // 2: stockpb.CheckIfItemsInStockResponse.Items:type_name -> orderpb.Item
	0, // 3: stockpb.StockService.GetItems:input_type -> stockpb.GetItemsRequest
	2, // 4: stockpb.StockService.CheckIfItemsInStock:input_type -> stockpb.CheckIfItemsInStockRequest
	1, // 5: stockpb.StockService.GetItems:output_type -> stockpb.GetItemsResponse
	3, // 6: stockpb.StockService.CheckIfItemsInStock:output_type -> stockpb.CheckIfItemsInStockResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_stockpb_stock_proto_init() }
func file_stockpb_stock_proto_init() {
	if File_stockpb_stock_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_stockpb_stock_proto_rawDesc), len(file_stockpb_stock_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stockpb_stock_proto_goTypes,
		DependencyIndexes: file_stockpb_stock_proto_depIdxs,
		MessageInfos:      file_stockpb_stock_proto_msgTypes,
	}.Build()
	File_stockpb_stock_proto = out.File
	file_stockpb_stock_proto_goTypes = nil
	file_stockpb_stock_proto_depIdxs = nil
}
