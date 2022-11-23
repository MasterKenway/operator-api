// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using PublishServiceMetadataRequest within kubernetes types, where deepcopy-gen is used.
func (in *PublishServiceMetadataRequest) DeepCopyInto(out *PublishServiceMetadataRequest) {
	p := proto.Clone(in).(*PublishServiceMetadataRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishServiceMetadataRequest. Required by controller-gen.
func (in *PublishServiceMetadataRequest) DeepCopy() *PublishServiceMetadataRequest {
	if in == nil {
		return nil
	}
	out := new(PublishServiceMetadataRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PublishServiceMetadataRequest. Required by controller-gen.
func (in *PublishServiceMetadataRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PublishServiceMetadataResponse within kubernetes types, where deepcopy-gen is used.
func (in *PublishServiceMetadataResponse) DeepCopyInto(out *PublishServiceMetadataResponse) {
	p := proto.Clone(in).(*PublishServiceMetadataResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishServiceMetadataResponse. Required by controller-gen.
func (in *PublishServiceMetadataResponse) DeepCopy() *PublishServiceMetadataResponse {
	if in == nil {
		return nil
	}
	out := new(PublishServiceMetadataResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PublishServiceMetadataResponse. Required by controller-gen.
func (in *PublishServiceMetadataResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetServiceMetadataRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetServiceMetadataRequest) DeepCopyInto(out *GetServiceMetadataRequest) {
	p := proto.Clone(in).(*GetServiceMetadataRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetServiceMetadataRequest. Required by controller-gen.
func (in *GetServiceMetadataRequest) DeepCopy() *GetServiceMetadataRequest {
	if in == nil {
		return nil
	}
	out := new(GetServiceMetadataRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetServiceMetadataRequest. Required by controller-gen.
func (in *GetServiceMetadataRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetServiceMetadataResponse within kubernetes types, where deepcopy-gen is used.
func (in *GetServiceMetadataResponse) DeepCopyInto(out *GetServiceMetadataResponse) {
	p := proto.Clone(in).(*GetServiceMetadataResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetServiceMetadataResponse. Required by controller-gen.
func (in *GetServiceMetadataResponse) DeepCopy() *GetServiceMetadataResponse {
	if in == nil {
		return nil
	}
	out := new(GetServiceMetadataResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetServiceMetadataResponse. Required by controller-gen.
func (in *GetServiceMetadataResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}