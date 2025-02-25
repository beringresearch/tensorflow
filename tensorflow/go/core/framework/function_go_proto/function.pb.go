// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: tensorflow/core/framework/function.proto

package function_go_proto

import (
	attr_value_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/attr_value_go_proto"
	node_def_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/node_def_go_proto"
	op_def_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/op_def_go_proto"
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

// A library is a set of named functions.
type FunctionDefLibrary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Function            []*FunctionDef        `protobuf:"bytes,1,rep,name=function,proto3" json:"function,omitempty"`
	Gradient            []*GradientDef        `protobuf:"bytes,2,rep,name=gradient,proto3" json:"gradient,omitempty"`
	RegisteredGradients []*RegisteredGradient `protobuf:"bytes,3,rep,name=registered_gradients,json=registeredGradients,proto3" json:"registered_gradients,omitempty"`
}

func (x *FunctionDefLibrary) Reset() {
	*x = FunctionDefLibrary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_function_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionDefLibrary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionDefLibrary) ProtoMessage() {}

func (x *FunctionDefLibrary) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_function_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionDefLibrary.ProtoReflect.Descriptor instead.
func (*FunctionDefLibrary) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_function_proto_rawDescGZIP(), []int{0}
}

func (x *FunctionDefLibrary) GetFunction() []*FunctionDef {
	if x != nil {
		return x.Function
	}
	return nil
}

func (x *FunctionDefLibrary) GetGradient() []*GradientDef {
	if x != nil {
		return x.Gradient
	}
	return nil
}

func (x *FunctionDefLibrary) GetRegisteredGradients() []*RegisteredGradient {
	if x != nil {
		return x.RegisteredGradients
	}
	return nil
}

// A function can be instantiated when the runtime can bind every attr
// with a value. When a GraphDef has a call to a function, it must
// have binding for every attr defined in the signature.
//
// TODO(zhifengc):
//   * device spec, etc.
type FunctionDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The definition of the function's name, arguments, return values,
	// attrs etc.
	Signature *op_def_go_proto.OpDef `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Attributes specific to this function definition.
	Attr    map[string]*attr_value_go_proto.AttrValue `protobuf:"bytes,5,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ArgAttr map[uint32]*FunctionDef_ArgAttrs          `protobuf:"bytes,7,rep,name=arg_attr,json=argAttr,proto3" json:"arg_attr,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Unique IDs for each resource argument, used to track aliasing resources. If
	// Argument A and Argument B alias each other, then
	// resource_arg_unique_ids[A.index] == resource_arg_unique_ids[B.index].
	//
	// If this field is empty, none of the arguments could alias; otherwise, every
	// resource argument should have an entry in this field.
	//
	// When instantiated, the unique IDs will be attached to the _Arg nodes'
	// "_resource_arg_unique_id" attribute.
	ResourceArgUniqueId map[uint32]uint32 `protobuf:"bytes,8,rep,name=resource_arg_unique_id,json=resourceArgUniqueId,proto3" json:"resource_arg_unique_id,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// By convention, "op" in node_def is resolved by consulting with a
	// user-defined library first. If not resolved, "func" is assumed to
	// be a builtin op.
	NodeDef []*node_def_go_proto.NodeDef `protobuf:"bytes,3,rep,name=node_def,json=nodeDef,proto3" json:"node_def,omitempty"`
	// A mapping from the output arg names from `signature` to the
	// outputs from `node_def` that should be returned by the function.
	Ret map[string]string `protobuf:"bytes,4,rep,name=ret,proto3" json:"ret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A mapping from control output names from `signature` to node names in
	// `node_def` which should be control outputs of this function.
	ControlRet map[string]string `protobuf:"bytes,6,rep,name=control_ret,json=controlRet,proto3" json:"control_ret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FunctionDef) Reset() {
	*x = FunctionDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_function_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionDef) ProtoMessage() {}

func (x *FunctionDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_function_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionDef.ProtoReflect.Descriptor instead.
func (*FunctionDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_function_proto_rawDescGZIP(), []int{1}
}

func (x *FunctionDef) GetSignature() *op_def_go_proto.OpDef {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *FunctionDef) GetAttr() map[string]*attr_value_go_proto.AttrValue {
	if x != nil {
		return x.Attr
	}
	return nil
}

func (x *FunctionDef) GetArgAttr() map[uint32]*FunctionDef_ArgAttrs {
	if x != nil {
		return x.ArgAttr
	}
	return nil
}

func (x *FunctionDef) GetResourceArgUniqueId() map[uint32]uint32 {
	if x != nil {
		return x.ResourceArgUniqueId
	}
	return nil
}

func (x *FunctionDef) GetNodeDef() []*node_def_go_proto.NodeDef {
	if x != nil {
		return x.NodeDef
	}
	return nil
}

func (x *FunctionDef) GetRet() map[string]string {
	if x != nil {
		return x.Ret
	}
	return nil
}

func (x *FunctionDef) GetControlRet() map[string]string {
	if x != nil {
		return x.ControlRet
	}
	return nil
}

// GradientDef defines the gradient function of a function defined in
// a function library.
//
// A gradient function g (specified by gradient_func) for a function f
// (specified by function_name) must follow the following:
//
// The function 'f' must be a numerical function which takes N inputs
// and produces M outputs. Its gradient function 'g', which is a
// function taking N + M inputs and produces N outputs.
//
// I.e. if we have
//    (y1, y2, ..., y_M) = f(x1, x2, ..., x_N),
// then, g is
//    (dL/dx1, dL/dx2, ..., dL/dx_N) = g(x1, x2, ..., x_N,
//                                      dL/dy1, dL/dy2, ..., dL/dy_M),
// where L is a scalar-value function of (x1, x2, ..., xN) (e.g., the
// loss function). dL/dx_i is the partial derivative of L with respect
// to x_i.
type GradientDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"` // The function name.
	GradientFunc string `protobuf:"bytes,2,opt,name=gradient_func,json=gradientFunc,proto3" json:"gradient_func,omitempty"` // The gradient function's name.
}

func (x *GradientDef) Reset() {
	*x = GradientDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_function_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GradientDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GradientDef) ProtoMessage() {}

func (x *GradientDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_function_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GradientDef.ProtoReflect.Descriptor instead.
func (*GradientDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_function_proto_rawDescGZIP(), []int{2}
}

func (x *GradientDef) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *GradientDef) GetGradientFunc() string {
	if x != nil {
		return x.GradientFunc
	}
	return ""
}

// RegisteredGradient stores a gradient function that is registered in the
// gradients library and used in the ops of a function in the function library.
// Unlike GradientDef, these gradients are identified by op type, and not
// directly linked to any function.
type RegisteredGradient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GradientFunc     string `protobuf:"bytes,1,opt,name=gradient_func,json=gradientFunc,proto3" json:"gradient_func,omitempty"`               // The gradient function's name.
	RegisteredOpType string `protobuf:"bytes,2,opt,name=registered_op_type,json=registeredOpType,proto3" json:"registered_op_type,omitempty"` // The gradient function's registered op type.
}

func (x *RegisteredGradient) Reset() {
	*x = RegisteredGradient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_function_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteredGradient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteredGradient) ProtoMessage() {}

func (x *RegisteredGradient) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_function_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteredGradient.ProtoReflect.Descriptor instead.
func (*RegisteredGradient) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_function_proto_rawDescGZIP(), []int{3}
}

func (x *RegisteredGradient) GetGradientFunc() string {
	if x != nil {
		return x.GradientFunc
	}
	return ""
}

func (x *RegisteredGradient) GetRegisteredOpType() string {
	if x != nil {
		return x.RegisteredOpType
	}
	return ""
}

// Attributes for function arguments. These attributes are the same set of
// valid attributes as to _Arg nodes.
type FunctionDef_ArgAttrs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attr map[string]*attr_value_go_proto.AttrValue `protobuf:"bytes,1,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FunctionDef_ArgAttrs) Reset() {
	*x = FunctionDef_ArgAttrs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_function_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionDef_ArgAttrs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionDef_ArgAttrs) ProtoMessage() {}

func (x *FunctionDef_ArgAttrs) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_function_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionDef_ArgAttrs.ProtoReflect.Descriptor instead.
func (*FunctionDef_ArgAttrs) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_function_proto_rawDescGZIP(), []int{1, 1}
}

func (x *FunctionDef_ArgAttrs) GetAttr() map[string]*attr_value_go_proto.AttrValue {
	if x != nil {
		return x.Attr
	}
	return nil
}

var File_tensorflow_core_framework_function_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_function_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x70, 0x5f, 0x64, 0x65, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x12, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x66, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x66, 0x52, 0x08, 0x67, 0x72, 0x61,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x51, 0x0a, 0x14, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x47, 0x72, 0x61, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x47,
	0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xdb, 0x07, 0x0a, 0x0b, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x12, 0x2f, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4f, 0x70, 0x44, 0x65, 0x66, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x61, 0x74, 0x74,
	0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72,
	0x12, 0x3f, 0x0a, 0x08, 0x61, 0x72, 0x67, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x2e, 0x41, 0x72, 0x67, 0x41,
	0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x72, 0x67, 0x41, 0x74, 0x74,
	0x72, 0x12, 0x65, 0x0a, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x72,
	0x67, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x41, 0x72, 0x67, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x13, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x72, 0x67,
	0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x64, 0x65, 0x66, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x52,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x12, 0x32, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x2e, 0x52,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x48, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x72, 0x65, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x52, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x52, 0x65, 0x74, 0x1a, 0x4e, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x9a, 0x01, 0x0a, 0x08, 0x41, 0x72, 0x67, 0x41, 0x74,
	0x74, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x2e, 0x41, 0x72, 0x67, 0x41, 0x74,
	0x74, 0x72, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61,
	0x74, 0x74, 0x72, 0x1a, 0x4e, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x5c, 0x0a, 0x0c, 0x41, 0x72, 0x67, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x2e, 0x41, 0x72,
	0x67, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x46, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x72, 0x67,
	0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x57, 0x0a, 0x0b, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x66, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x72,
	0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x22,
	0x67, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x47, 0x72, 0x61,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x72,
	0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x42, 0x80, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_function_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_function_proto_rawDescData = file_tensorflow_core_framework_function_proto_rawDesc
)

func file_tensorflow_core_framework_function_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_function_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_function_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_function_proto_rawDescData)
	})
	return file_tensorflow_core_framework_function_proto_rawDescData
}

var file_tensorflow_core_framework_function_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_tensorflow_core_framework_function_proto_goTypes = []interface{}{
	(*FunctionDefLibrary)(nil),            // 0: tensorflow.FunctionDefLibrary
	(*FunctionDef)(nil),                   // 1: tensorflow.FunctionDef
	(*GradientDef)(nil),                   // 2: tensorflow.GradientDef
	(*RegisteredGradient)(nil),            // 3: tensorflow.RegisteredGradient
	nil,                                   // 4: tensorflow.FunctionDef.AttrEntry
	(*FunctionDef_ArgAttrs)(nil),          // 5: tensorflow.FunctionDef.ArgAttrs
	nil,                                   // 6: tensorflow.FunctionDef.ArgAttrEntry
	nil,                                   // 7: tensorflow.FunctionDef.ResourceArgUniqueIdEntry
	nil,                                   // 8: tensorflow.FunctionDef.RetEntry
	nil,                                   // 9: tensorflow.FunctionDef.ControlRetEntry
	nil,                                   // 10: tensorflow.FunctionDef.ArgAttrs.AttrEntry
	(*op_def_go_proto.OpDef)(nil),         // 11: tensorflow.OpDef
	(*node_def_go_proto.NodeDef)(nil),     // 12: tensorflow.NodeDef
	(*attr_value_go_proto.AttrValue)(nil), // 13: tensorflow.AttrValue
}
var file_tensorflow_core_framework_function_proto_depIdxs = []int32{
	1,  // 0: tensorflow.FunctionDefLibrary.function:type_name -> tensorflow.FunctionDef
	2,  // 1: tensorflow.FunctionDefLibrary.gradient:type_name -> tensorflow.GradientDef
	3,  // 2: tensorflow.FunctionDefLibrary.registered_gradients:type_name -> tensorflow.RegisteredGradient
	11, // 3: tensorflow.FunctionDef.signature:type_name -> tensorflow.OpDef
	4,  // 4: tensorflow.FunctionDef.attr:type_name -> tensorflow.FunctionDef.AttrEntry
	6,  // 5: tensorflow.FunctionDef.arg_attr:type_name -> tensorflow.FunctionDef.ArgAttrEntry
	7,  // 6: tensorflow.FunctionDef.resource_arg_unique_id:type_name -> tensorflow.FunctionDef.ResourceArgUniqueIdEntry
	12, // 7: tensorflow.FunctionDef.node_def:type_name -> tensorflow.NodeDef
	8,  // 8: tensorflow.FunctionDef.ret:type_name -> tensorflow.FunctionDef.RetEntry
	9,  // 9: tensorflow.FunctionDef.control_ret:type_name -> tensorflow.FunctionDef.ControlRetEntry
	13, // 10: tensorflow.FunctionDef.AttrEntry.value:type_name -> tensorflow.AttrValue
	10, // 11: tensorflow.FunctionDef.ArgAttrs.attr:type_name -> tensorflow.FunctionDef.ArgAttrs.AttrEntry
	5,  // 12: tensorflow.FunctionDef.ArgAttrEntry.value:type_name -> tensorflow.FunctionDef.ArgAttrs
	13, // 13: tensorflow.FunctionDef.ArgAttrs.AttrEntry.value:type_name -> tensorflow.AttrValue
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_function_proto_init() }
func file_tensorflow_core_framework_function_proto_init() {
	if File_tensorflow_core_framework_function_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_function_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionDefLibrary); i {
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
		file_tensorflow_core_framework_function_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionDef); i {
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
		file_tensorflow_core_framework_function_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GradientDef); i {
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
		file_tensorflow_core_framework_function_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisteredGradient); i {
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
		file_tensorflow_core_framework_function_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionDef_ArgAttrs); i {
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
			RawDescriptor: file_tensorflow_core_framework_function_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_function_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_function_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_function_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_function_proto = out.File
	file_tensorflow_core_framework_function_proto_rawDesc = nil
	file_tensorflow_core_framework_function_proto_goTypes = nil
	file_tensorflow_core_framework_function_proto_depIdxs = nil
}
