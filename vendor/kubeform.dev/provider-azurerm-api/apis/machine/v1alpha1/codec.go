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
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecIdentity{}).Type1()):      LearningComputeClusterSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecScaleSettings{}).Type1()): LearningComputeClusterSpecScaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecSsh{}).Type1()):           LearningComputeClusterSpecSshCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecAssignToUser{}).Type1()): LearningComputeInstanceSpecAssignToUserCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecIdentity{}).Type1()):     LearningComputeInstanceSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecSsh{}).Type1()):          LearningComputeInstanceSpecSshCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecIdentity{}).Type1()):    LearningInferenceClusterSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecSsl{}).Type1()):         LearningInferenceClusterSpecSslCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningSynapseSparkSpecIdentity{}).Type1()):        LearningSynapseSparkSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecEncryption{}).Type1()):         LearningWorkspaceSpecEncryptionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecIdentity{}).Type1()):           LearningWorkspaceSpecIdentityCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecIdentity{}).Type1()):      LearningComputeClusterSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecScaleSettings{}).Type1()): LearningComputeClusterSpecScaleSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecSsh{}).Type1()):           LearningComputeClusterSpecSshCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecAssignToUser{}).Type1()): LearningComputeInstanceSpecAssignToUserCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecIdentity{}).Type1()):     LearningComputeInstanceSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecSsh{}).Type1()):          LearningComputeInstanceSpecSshCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecIdentity{}).Type1()):    LearningInferenceClusterSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecSsl{}).Type1()):         LearningInferenceClusterSpecSslCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningSynapseSparkSpecIdentity{}).Type1()):        LearningSynapseSparkSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecEncryption{}).Type1()):         LearningWorkspaceSpecEncryptionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecIdentity{}).Type1()):           LearningWorkspaceSpecIdentityCodec{},
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
type LearningComputeClusterSpecIdentityCodec struct {
}

func (LearningComputeClusterSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningComputeClusterSpecIdentity)(ptr) == nil
}

func (LearningComputeClusterSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningComputeClusterSpecIdentity)(ptr)
	var objs []LearningComputeClusterSpecIdentity
	if obj != nil {
		objs = []LearningComputeClusterSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningComputeClusterSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningComputeClusterSpecIdentity)(ptr) = LearningComputeClusterSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningComputeClusterSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningComputeClusterSpecIdentity)(ptr) = objs[0]
			} else {
				*(*LearningComputeClusterSpecIdentity)(ptr) = LearningComputeClusterSpecIdentity{}
			}
		} else {
			*(*LearningComputeClusterSpecIdentity)(ptr) = LearningComputeClusterSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningComputeClusterSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningComputeClusterSpecIdentity)(ptr) = obj
		} else {
			*(*LearningComputeClusterSpecIdentity)(ptr) = LearningComputeClusterSpecIdentity{}
		}
	default:
		iter.ReportError("decode LearningComputeClusterSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningComputeClusterSpecScaleSettingsCodec struct {
}

func (LearningComputeClusterSpecScaleSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningComputeClusterSpecScaleSettings)(ptr) == nil
}

func (LearningComputeClusterSpecScaleSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningComputeClusterSpecScaleSettings)(ptr)
	var objs []LearningComputeClusterSpecScaleSettings
	if obj != nil {
		objs = []LearningComputeClusterSpecScaleSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecScaleSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningComputeClusterSpecScaleSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningComputeClusterSpecScaleSettings)(ptr) = LearningComputeClusterSpecScaleSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningComputeClusterSpecScaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecScaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningComputeClusterSpecScaleSettings)(ptr) = objs[0]
			} else {
				*(*LearningComputeClusterSpecScaleSettings)(ptr) = LearningComputeClusterSpecScaleSettings{}
			}
		} else {
			*(*LearningComputeClusterSpecScaleSettings)(ptr) = LearningComputeClusterSpecScaleSettings{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningComputeClusterSpecScaleSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecScaleSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningComputeClusterSpecScaleSettings)(ptr) = obj
		} else {
			*(*LearningComputeClusterSpecScaleSettings)(ptr) = LearningComputeClusterSpecScaleSettings{}
		}
	default:
		iter.ReportError("decode LearningComputeClusterSpecScaleSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningComputeClusterSpecSshCodec struct {
}

func (LearningComputeClusterSpecSshCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningComputeClusterSpecSsh)(ptr) == nil
}

func (LearningComputeClusterSpecSshCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningComputeClusterSpecSsh)(ptr)
	var objs []LearningComputeClusterSpecSsh
	if obj != nil {
		objs = []LearningComputeClusterSpecSsh{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecSsh{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningComputeClusterSpecSshCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningComputeClusterSpecSsh)(ptr) = LearningComputeClusterSpecSsh{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningComputeClusterSpecSsh

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecSsh{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningComputeClusterSpecSsh)(ptr) = objs[0]
			} else {
				*(*LearningComputeClusterSpecSsh)(ptr) = LearningComputeClusterSpecSsh{}
			}
		} else {
			*(*LearningComputeClusterSpecSsh)(ptr) = LearningComputeClusterSpecSsh{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningComputeClusterSpecSsh

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeClusterSpecSsh{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningComputeClusterSpecSsh)(ptr) = obj
		} else {
			*(*LearningComputeClusterSpecSsh)(ptr) = LearningComputeClusterSpecSsh{}
		}
	default:
		iter.ReportError("decode LearningComputeClusterSpecSsh", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningComputeInstanceSpecAssignToUserCodec struct {
}

func (LearningComputeInstanceSpecAssignToUserCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningComputeInstanceSpecAssignToUser)(ptr) == nil
}

func (LearningComputeInstanceSpecAssignToUserCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningComputeInstanceSpecAssignToUser)(ptr)
	var objs []LearningComputeInstanceSpecAssignToUser
	if obj != nil {
		objs = []LearningComputeInstanceSpecAssignToUser{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecAssignToUser{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningComputeInstanceSpecAssignToUserCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningComputeInstanceSpecAssignToUser)(ptr) = LearningComputeInstanceSpecAssignToUser{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningComputeInstanceSpecAssignToUser

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecAssignToUser{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningComputeInstanceSpecAssignToUser)(ptr) = objs[0]
			} else {
				*(*LearningComputeInstanceSpecAssignToUser)(ptr) = LearningComputeInstanceSpecAssignToUser{}
			}
		} else {
			*(*LearningComputeInstanceSpecAssignToUser)(ptr) = LearningComputeInstanceSpecAssignToUser{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningComputeInstanceSpecAssignToUser

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecAssignToUser{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningComputeInstanceSpecAssignToUser)(ptr) = obj
		} else {
			*(*LearningComputeInstanceSpecAssignToUser)(ptr) = LearningComputeInstanceSpecAssignToUser{}
		}
	default:
		iter.ReportError("decode LearningComputeInstanceSpecAssignToUser", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningComputeInstanceSpecIdentityCodec struct {
}

func (LearningComputeInstanceSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningComputeInstanceSpecIdentity)(ptr) == nil
}

func (LearningComputeInstanceSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningComputeInstanceSpecIdentity)(ptr)
	var objs []LearningComputeInstanceSpecIdentity
	if obj != nil {
		objs = []LearningComputeInstanceSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningComputeInstanceSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningComputeInstanceSpecIdentity)(ptr) = LearningComputeInstanceSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningComputeInstanceSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningComputeInstanceSpecIdentity)(ptr) = objs[0]
			} else {
				*(*LearningComputeInstanceSpecIdentity)(ptr) = LearningComputeInstanceSpecIdentity{}
			}
		} else {
			*(*LearningComputeInstanceSpecIdentity)(ptr) = LearningComputeInstanceSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningComputeInstanceSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningComputeInstanceSpecIdentity)(ptr) = obj
		} else {
			*(*LearningComputeInstanceSpecIdentity)(ptr) = LearningComputeInstanceSpecIdentity{}
		}
	default:
		iter.ReportError("decode LearningComputeInstanceSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningComputeInstanceSpecSshCodec struct {
}

func (LearningComputeInstanceSpecSshCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningComputeInstanceSpecSsh)(ptr) == nil
}

func (LearningComputeInstanceSpecSshCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningComputeInstanceSpecSsh)(ptr)
	var objs []LearningComputeInstanceSpecSsh
	if obj != nil {
		objs = []LearningComputeInstanceSpecSsh{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecSsh{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningComputeInstanceSpecSshCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningComputeInstanceSpecSsh)(ptr) = LearningComputeInstanceSpecSsh{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningComputeInstanceSpecSsh

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecSsh{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningComputeInstanceSpecSsh)(ptr) = objs[0]
			} else {
				*(*LearningComputeInstanceSpecSsh)(ptr) = LearningComputeInstanceSpecSsh{}
			}
		} else {
			*(*LearningComputeInstanceSpecSsh)(ptr) = LearningComputeInstanceSpecSsh{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningComputeInstanceSpecSsh

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningComputeInstanceSpecSsh{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningComputeInstanceSpecSsh)(ptr) = obj
		} else {
			*(*LearningComputeInstanceSpecSsh)(ptr) = LearningComputeInstanceSpecSsh{}
		}
	default:
		iter.ReportError("decode LearningComputeInstanceSpecSsh", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningInferenceClusterSpecIdentityCodec struct {
}

func (LearningInferenceClusterSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningInferenceClusterSpecIdentity)(ptr) == nil
}

func (LearningInferenceClusterSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningInferenceClusterSpecIdentity)(ptr)
	var objs []LearningInferenceClusterSpecIdentity
	if obj != nil {
		objs = []LearningInferenceClusterSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningInferenceClusterSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningInferenceClusterSpecIdentity)(ptr) = LearningInferenceClusterSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningInferenceClusterSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningInferenceClusterSpecIdentity)(ptr) = objs[0]
			} else {
				*(*LearningInferenceClusterSpecIdentity)(ptr) = LearningInferenceClusterSpecIdentity{}
			}
		} else {
			*(*LearningInferenceClusterSpecIdentity)(ptr) = LearningInferenceClusterSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningInferenceClusterSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningInferenceClusterSpecIdentity)(ptr) = obj
		} else {
			*(*LearningInferenceClusterSpecIdentity)(ptr) = LearningInferenceClusterSpecIdentity{}
		}
	default:
		iter.ReportError("decode LearningInferenceClusterSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningInferenceClusterSpecSslCodec struct {
}

func (LearningInferenceClusterSpecSslCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningInferenceClusterSpecSsl)(ptr) == nil
}

func (LearningInferenceClusterSpecSslCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningInferenceClusterSpecSsl)(ptr)
	var objs []LearningInferenceClusterSpecSsl
	if obj != nil {
		objs = []LearningInferenceClusterSpecSsl{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecSsl{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningInferenceClusterSpecSslCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningInferenceClusterSpecSsl)(ptr) = LearningInferenceClusterSpecSsl{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningInferenceClusterSpecSsl

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecSsl{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningInferenceClusterSpecSsl)(ptr) = objs[0]
			} else {
				*(*LearningInferenceClusterSpecSsl)(ptr) = LearningInferenceClusterSpecSsl{}
			}
		} else {
			*(*LearningInferenceClusterSpecSsl)(ptr) = LearningInferenceClusterSpecSsl{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningInferenceClusterSpecSsl

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningInferenceClusterSpecSsl{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningInferenceClusterSpecSsl)(ptr) = obj
		} else {
			*(*LearningInferenceClusterSpecSsl)(ptr) = LearningInferenceClusterSpecSsl{}
		}
	default:
		iter.ReportError("decode LearningInferenceClusterSpecSsl", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningSynapseSparkSpecIdentityCodec struct {
}

func (LearningSynapseSparkSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningSynapseSparkSpecIdentity)(ptr) == nil
}

func (LearningSynapseSparkSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningSynapseSparkSpecIdentity)(ptr)
	var objs []LearningSynapseSparkSpecIdentity
	if obj != nil {
		objs = []LearningSynapseSparkSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningSynapseSparkSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningSynapseSparkSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningSynapseSparkSpecIdentity)(ptr) = LearningSynapseSparkSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningSynapseSparkSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningSynapseSparkSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningSynapseSparkSpecIdentity)(ptr) = objs[0]
			} else {
				*(*LearningSynapseSparkSpecIdentity)(ptr) = LearningSynapseSparkSpecIdentity{}
			}
		} else {
			*(*LearningSynapseSparkSpecIdentity)(ptr) = LearningSynapseSparkSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningSynapseSparkSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningSynapseSparkSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningSynapseSparkSpecIdentity)(ptr) = obj
		} else {
			*(*LearningSynapseSparkSpecIdentity)(ptr) = LearningSynapseSparkSpecIdentity{}
		}
	default:
		iter.ReportError("decode LearningSynapseSparkSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningWorkspaceSpecEncryptionCodec struct {
}

func (LearningWorkspaceSpecEncryptionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningWorkspaceSpecEncryption)(ptr) == nil
}

func (LearningWorkspaceSpecEncryptionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningWorkspaceSpecEncryption)(ptr)
	var objs []LearningWorkspaceSpecEncryption
	if obj != nil {
		objs = []LearningWorkspaceSpecEncryption{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecEncryption{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningWorkspaceSpecEncryptionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningWorkspaceSpecEncryption)(ptr) = LearningWorkspaceSpecEncryption{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningWorkspaceSpecEncryption

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecEncryption{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningWorkspaceSpecEncryption)(ptr) = objs[0]
			} else {
				*(*LearningWorkspaceSpecEncryption)(ptr) = LearningWorkspaceSpecEncryption{}
			}
		} else {
			*(*LearningWorkspaceSpecEncryption)(ptr) = LearningWorkspaceSpecEncryption{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningWorkspaceSpecEncryption

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecEncryption{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningWorkspaceSpecEncryption)(ptr) = obj
		} else {
			*(*LearningWorkspaceSpecEncryption)(ptr) = LearningWorkspaceSpecEncryption{}
		}
	default:
		iter.ReportError("decode LearningWorkspaceSpecEncryption", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LearningWorkspaceSpecIdentityCodec struct {
}

func (LearningWorkspaceSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LearningWorkspaceSpecIdentity)(ptr) == nil
}

func (LearningWorkspaceSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LearningWorkspaceSpecIdentity)(ptr)
	var objs []LearningWorkspaceSpecIdentity
	if obj != nil {
		objs = []LearningWorkspaceSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LearningWorkspaceSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LearningWorkspaceSpecIdentity)(ptr) = LearningWorkspaceSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LearningWorkspaceSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LearningWorkspaceSpecIdentity)(ptr) = objs[0]
			} else {
				*(*LearningWorkspaceSpecIdentity)(ptr) = LearningWorkspaceSpecIdentity{}
			}
		} else {
			*(*LearningWorkspaceSpecIdentity)(ptr) = LearningWorkspaceSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LearningWorkspaceSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LearningWorkspaceSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LearningWorkspaceSpecIdentity)(ptr) = obj
		} else {
			*(*LearningWorkspaceSpecIdentity)(ptr) = LearningWorkspaceSpecIdentity{}
		}
	default:
		iter.ReportError("decode LearningWorkspaceSpecIdentity", "unexpected JSON type")
	}
}
