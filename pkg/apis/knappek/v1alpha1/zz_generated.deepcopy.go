// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	mongodbatlas "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasCluster) DeepCopyInto(out *MongoDBAtlasCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasCluster.
func (in *MongoDBAtlasCluster) DeepCopy() *MongoDBAtlasCluster {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBAtlasCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasClusterList) DeepCopyInto(out *MongoDBAtlasClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDBAtlasCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasClusterList.
func (in *MongoDBAtlasClusterList) DeepCopy() *MongoDBAtlasClusterList {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBAtlasClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasClusterRequestBody) DeepCopyInto(out *MongoDBAtlasClusterRequestBody) {
	*out = *in
	if in.ReplicationSpec != nil {
		in, out := &in.ReplicationSpec, &out.ReplicationSpec
		*out = make(map[string]mongodbatlas.ReplicationSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.AutoScaling = in.AutoScaling
	out.ProviderSettings = in.ProviderSettings
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasClusterRequestBody.
func (in *MongoDBAtlasClusterRequestBody) DeepCopy() *MongoDBAtlasClusterRequestBody {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasClusterRequestBody)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasClusterSpec) DeepCopyInto(out *MongoDBAtlasClusterSpec) {
	*out = *in
	in.MongoDBAtlasClusterRequestBody.DeepCopyInto(&out.MongoDBAtlasClusterRequestBody)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasClusterSpec.
func (in *MongoDBAtlasClusterSpec) DeepCopy() *MongoDBAtlasClusterSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasClusterStatus) DeepCopyInto(out *MongoDBAtlasClusterStatus) {
	*out = *in
	in.MongoDBAtlasClusterRequestBody.DeepCopyInto(&out.MongoDBAtlasClusterRequestBody)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasClusterStatus.
func (in *MongoDBAtlasClusterStatus) DeepCopy() *MongoDBAtlasClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasDatabaseUser) DeepCopyInto(out *MongoDBAtlasDatabaseUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasDatabaseUser.
func (in *MongoDBAtlasDatabaseUser) DeepCopy() *MongoDBAtlasDatabaseUser {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasDatabaseUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBAtlasDatabaseUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasDatabaseUserList) DeepCopyInto(out *MongoDBAtlasDatabaseUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDBAtlasDatabaseUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasDatabaseUserList.
func (in *MongoDBAtlasDatabaseUserList) DeepCopy() *MongoDBAtlasDatabaseUserList {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasDatabaseUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBAtlasDatabaseUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasDatabaseUserSpec) DeepCopyInto(out *MongoDBAtlasDatabaseUserSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasDatabaseUserSpec.
func (in *MongoDBAtlasDatabaseUserSpec) DeepCopy() *MongoDBAtlasDatabaseUserSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasDatabaseUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasDatabaseUserStatus) DeepCopyInto(out *MongoDBAtlasDatabaseUserStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasDatabaseUserStatus.
func (in *MongoDBAtlasDatabaseUserStatus) DeepCopy() *MongoDBAtlasDatabaseUserStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasDatabaseUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasProject) DeepCopyInto(out *MongoDBAtlasProject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasProject.
func (in *MongoDBAtlasProject) DeepCopy() *MongoDBAtlasProject {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBAtlasProject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasProjectList) DeepCopyInto(out *MongoDBAtlasProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDBAtlasProject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasProjectList.
func (in *MongoDBAtlasProjectList) DeepCopy() *MongoDBAtlasProjectList {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBAtlasProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasProjectSpec) DeepCopyInto(out *MongoDBAtlasProjectSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasProjectSpec.
func (in *MongoDBAtlasProjectSpec) DeepCopy() *MongoDBAtlasProjectSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasProjectStatus) DeepCopyInto(out *MongoDBAtlasProjectStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasProjectStatus.
func (in *MongoDBAtlasProjectStatus) DeepCopy() *MongoDBAtlasProjectStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasProjectStatus)
	in.DeepCopyInto(out)
	return out
}
