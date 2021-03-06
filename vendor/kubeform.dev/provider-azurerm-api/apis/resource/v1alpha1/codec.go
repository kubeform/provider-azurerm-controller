/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(GroupCostManagementExportSpecExportDataOptions{}).Type1()):         GroupCostManagementExportSpecExportDataOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GroupCostManagementExportSpecExportDataStorageLocation{}).Type1()): GroupCostManagementExportSpecExportDataStorageLocationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GroupPolicyAssignmentSpecIdentity{}).Type1()):                      GroupPolicyAssignmentSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicyAssignmentSpecIdentity{}).Type1()):                           PolicyAssignmentSpecIdentityCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(GroupCostManagementExportSpecExportDataOptions{}).Type1()):         GroupCostManagementExportSpecExportDataOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GroupCostManagementExportSpecExportDataStorageLocation{}).Type1()): GroupCostManagementExportSpecExportDataStorageLocationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GroupPolicyAssignmentSpecIdentity{}).Type1()):                      GroupPolicyAssignmentSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicyAssignmentSpecIdentity{}).Type1()):                           PolicyAssignmentSpecIdentityCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type GroupCostManagementExportSpecExportDataOptionsCodec struct {
}

func (GroupCostManagementExportSpecExportDataOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GroupCostManagementExportSpecExportDataOptions)(ptr) == nil
}

func (GroupCostManagementExportSpecExportDataOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GroupCostManagementExportSpecExportDataOptions)(ptr)
	var objs []GroupCostManagementExportSpecExportDataOptions
	if obj != nil {
		objs = []GroupCostManagementExportSpecExportDataOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GroupCostManagementExportSpecExportDataOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GroupCostManagementExportSpecExportDataOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GroupCostManagementExportSpecExportDataOptions)(ptr) = GroupCostManagementExportSpecExportDataOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GroupCostManagementExportSpecExportDataOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GroupCostManagementExportSpecExportDataOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GroupCostManagementExportSpecExportDataOptions)(ptr) = objs[0]
			} else {
				*(*GroupCostManagementExportSpecExportDataOptions)(ptr) = GroupCostManagementExportSpecExportDataOptions{}
			}
		} else {
			*(*GroupCostManagementExportSpecExportDataOptions)(ptr) = GroupCostManagementExportSpecExportDataOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GroupCostManagementExportSpecExportDataOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GroupCostManagementExportSpecExportDataOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GroupCostManagementExportSpecExportDataOptions)(ptr) = obj
		} else {
			*(*GroupCostManagementExportSpecExportDataOptions)(ptr) = GroupCostManagementExportSpecExportDataOptions{}
		}
	default:
		iter.ReportError("decode GroupCostManagementExportSpecExportDataOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GroupCostManagementExportSpecExportDataStorageLocationCodec struct {
}

func (GroupCostManagementExportSpecExportDataStorageLocationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GroupCostManagementExportSpecExportDataStorageLocation)(ptr) == nil
}

func (GroupCostManagementExportSpecExportDataStorageLocationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GroupCostManagementExportSpecExportDataStorageLocation)(ptr)
	var objs []GroupCostManagementExportSpecExportDataStorageLocation
	if obj != nil {
		objs = []GroupCostManagementExportSpecExportDataStorageLocation{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GroupCostManagementExportSpecExportDataStorageLocation{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GroupCostManagementExportSpecExportDataStorageLocationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GroupCostManagementExportSpecExportDataStorageLocation)(ptr) = GroupCostManagementExportSpecExportDataStorageLocation{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GroupCostManagementExportSpecExportDataStorageLocation

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GroupCostManagementExportSpecExportDataStorageLocation{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GroupCostManagementExportSpecExportDataStorageLocation)(ptr) = objs[0]
			} else {
				*(*GroupCostManagementExportSpecExportDataStorageLocation)(ptr) = GroupCostManagementExportSpecExportDataStorageLocation{}
			}
		} else {
			*(*GroupCostManagementExportSpecExportDataStorageLocation)(ptr) = GroupCostManagementExportSpecExportDataStorageLocation{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GroupCostManagementExportSpecExportDataStorageLocation

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GroupCostManagementExportSpecExportDataStorageLocation{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GroupCostManagementExportSpecExportDataStorageLocation)(ptr) = obj
		} else {
			*(*GroupCostManagementExportSpecExportDataStorageLocation)(ptr) = GroupCostManagementExportSpecExportDataStorageLocation{}
		}
	default:
		iter.ReportError("decode GroupCostManagementExportSpecExportDataStorageLocation", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GroupPolicyAssignmentSpecIdentityCodec struct {
}

func (GroupPolicyAssignmentSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GroupPolicyAssignmentSpecIdentity)(ptr) == nil
}

func (GroupPolicyAssignmentSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GroupPolicyAssignmentSpecIdentity)(ptr)
	var objs []GroupPolicyAssignmentSpecIdentity
	if obj != nil {
		objs = []GroupPolicyAssignmentSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GroupPolicyAssignmentSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GroupPolicyAssignmentSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GroupPolicyAssignmentSpecIdentity)(ptr) = GroupPolicyAssignmentSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GroupPolicyAssignmentSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GroupPolicyAssignmentSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GroupPolicyAssignmentSpecIdentity)(ptr) = objs[0]
			} else {
				*(*GroupPolicyAssignmentSpecIdentity)(ptr) = GroupPolicyAssignmentSpecIdentity{}
			}
		} else {
			*(*GroupPolicyAssignmentSpecIdentity)(ptr) = GroupPolicyAssignmentSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GroupPolicyAssignmentSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GroupPolicyAssignmentSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GroupPolicyAssignmentSpecIdentity)(ptr) = obj
		} else {
			*(*GroupPolicyAssignmentSpecIdentity)(ptr) = GroupPolicyAssignmentSpecIdentity{}
		}
	default:
		iter.ReportError("decode GroupPolicyAssignmentSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicyAssignmentSpecIdentityCodec struct {
}

func (PolicyAssignmentSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicyAssignmentSpecIdentity)(ptr) == nil
}

func (PolicyAssignmentSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicyAssignmentSpecIdentity)(ptr)
	var objs []PolicyAssignmentSpecIdentity
	if obj != nil {
		objs = []PolicyAssignmentSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicyAssignmentSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicyAssignmentSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicyAssignmentSpecIdentity)(ptr) = PolicyAssignmentSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicyAssignmentSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicyAssignmentSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicyAssignmentSpecIdentity)(ptr) = objs[0]
			} else {
				*(*PolicyAssignmentSpecIdentity)(ptr) = PolicyAssignmentSpecIdentity{}
			}
		} else {
			*(*PolicyAssignmentSpecIdentity)(ptr) = PolicyAssignmentSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PolicyAssignmentSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicyAssignmentSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PolicyAssignmentSpecIdentity)(ptr) = obj
		} else {
			*(*PolicyAssignmentSpecIdentity)(ptr) = PolicyAssignmentSpecIdentity{}
		}
	default:
		iter.ReportError("decode PolicyAssignmentSpecIdentity", "unexpected JSON type")
	}
}
