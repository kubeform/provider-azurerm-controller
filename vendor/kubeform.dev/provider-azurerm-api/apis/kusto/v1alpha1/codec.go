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
		jsoniter.MustGetKind(reflect2.TypeOf(AttachedDatabaseConfigurationSpecSharing{}).Type1()): AttachedDatabaseConfigurationSpecSharingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecIdentity{}).Type1()):                      ClusterSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecOptimizedAutoScale{}).Type1()):            ClusterSpecOptimizedAutoScaleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecSku{}).Type1()):                           ClusterSpecSkuCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecVirtualNetworkConfiguration{}).Type1()):   ClusterSpecVirtualNetworkConfigurationCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(AttachedDatabaseConfigurationSpecSharing{}).Type1()): AttachedDatabaseConfigurationSpecSharingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecIdentity{}).Type1()):                      ClusterSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecOptimizedAutoScale{}).Type1()):            ClusterSpecOptimizedAutoScaleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecSku{}).Type1()):                           ClusterSpecSkuCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecVirtualNetworkConfiguration{}).Type1()):   ClusterSpecVirtualNetworkConfigurationCodec{},
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
type AttachedDatabaseConfigurationSpecSharingCodec struct {
}

func (AttachedDatabaseConfigurationSpecSharingCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AttachedDatabaseConfigurationSpecSharing)(ptr) == nil
}

func (AttachedDatabaseConfigurationSpecSharingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AttachedDatabaseConfigurationSpecSharing)(ptr)
	var objs []AttachedDatabaseConfigurationSpecSharing
	if obj != nil {
		objs = []AttachedDatabaseConfigurationSpecSharing{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttachedDatabaseConfigurationSpecSharing{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AttachedDatabaseConfigurationSpecSharingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AttachedDatabaseConfigurationSpecSharing)(ptr) = AttachedDatabaseConfigurationSpecSharing{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AttachedDatabaseConfigurationSpecSharing

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttachedDatabaseConfigurationSpecSharing{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AttachedDatabaseConfigurationSpecSharing)(ptr) = objs[0]
			} else {
				*(*AttachedDatabaseConfigurationSpecSharing)(ptr) = AttachedDatabaseConfigurationSpecSharing{}
			}
		} else {
			*(*AttachedDatabaseConfigurationSpecSharing)(ptr) = AttachedDatabaseConfigurationSpecSharing{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AttachedDatabaseConfigurationSpecSharing

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttachedDatabaseConfigurationSpecSharing{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AttachedDatabaseConfigurationSpecSharing)(ptr) = obj
		} else {
			*(*AttachedDatabaseConfigurationSpecSharing)(ptr) = AttachedDatabaseConfigurationSpecSharing{}
		}
	default:
		iter.ReportError("decode AttachedDatabaseConfigurationSpecSharing", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecIdentityCodec struct {
}

func (ClusterSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecIdentity)(ptr) == nil
}

func (ClusterSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecIdentity)(ptr)
	var objs []ClusterSpecIdentity
	if obj != nil {
		objs = []ClusterSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecIdentity)(ptr) = ClusterSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecIdentity)(ptr) = objs[0]
			} else {
				*(*ClusterSpecIdentity)(ptr) = ClusterSpecIdentity{}
			}
		} else {
			*(*ClusterSpecIdentity)(ptr) = ClusterSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecIdentity)(ptr) = obj
		} else {
			*(*ClusterSpecIdentity)(ptr) = ClusterSpecIdentity{}
		}
	default:
		iter.ReportError("decode ClusterSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecOptimizedAutoScaleCodec struct {
}

func (ClusterSpecOptimizedAutoScaleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecOptimizedAutoScale)(ptr) == nil
}

func (ClusterSpecOptimizedAutoScaleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecOptimizedAutoScale)(ptr)
	var objs []ClusterSpecOptimizedAutoScale
	if obj != nil {
		objs = []ClusterSpecOptimizedAutoScale{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecOptimizedAutoScale{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecOptimizedAutoScaleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecOptimizedAutoScale)(ptr) = ClusterSpecOptimizedAutoScale{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecOptimizedAutoScale

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecOptimizedAutoScale{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecOptimizedAutoScale)(ptr) = objs[0]
			} else {
				*(*ClusterSpecOptimizedAutoScale)(ptr) = ClusterSpecOptimizedAutoScale{}
			}
		} else {
			*(*ClusterSpecOptimizedAutoScale)(ptr) = ClusterSpecOptimizedAutoScale{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecOptimizedAutoScale

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecOptimizedAutoScale{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecOptimizedAutoScale)(ptr) = obj
		} else {
			*(*ClusterSpecOptimizedAutoScale)(ptr) = ClusterSpecOptimizedAutoScale{}
		}
	default:
		iter.ReportError("decode ClusterSpecOptimizedAutoScale", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecSkuCodec struct {
}

func (ClusterSpecSkuCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecSku)(ptr) == nil
}

func (ClusterSpecSkuCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecSku)(ptr)
	var objs []ClusterSpecSku
	if obj != nil {
		objs = []ClusterSpecSku{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecSku{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecSkuCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecSku)(ptr) = ClusterSpecSku{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecSku

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecSku{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecSku)(ptr) = objs[0]
			} else {
				*(*ClusterSpecSku)(ptr) = ClusterSpecSku{}
			}
		} else {
			*(*ClusterSpecSku)(ptr) = ClusterSpecSku{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecSku

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecSku{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecSku)(ptr) = obj
		} else {
			*(*ClusterSpecSku)(ptr) = ClusterSpecSku{}
		}
	default:
		iter.ReportError("decode ClusterSpecSku", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecVirtualNetworkConfigurationCodec struct {
}

func (ClusterSpecVirtualNetworkConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecVirtualNetworkConfiguration)(ptr) == nil
}

func (ClusterSpecVirtualNetworkConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecVirtualNetworkConfiguration)(ptr)
	var objs []ClusterSpecVirtualNetworkConfiguration
	if obj != nil {
		objs = []ClusterSpecVirtualNetworkConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecVirtualNetworkConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecVirtualNetworkConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecVirtualNetworkConfiguration)(ptr) = ClusterSpecVirtualNetworkConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecVirtualNetworkConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecVirtualNetworkConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecVirtualNetworkConfiguration)(ptr) = objs[0]
			} else {
				*(*ClusterSpecVirtualNetworkConfiguration)(ptr) = ClusterSpecVirtualNetworkConfiguration{}
			}
		} else {
			*(*ClusterSpecVirtualNetworkConfiguration)(ptr) = ClusterSpecVirtualNetworkConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecVirtualNetworkConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecVirtualNetworkConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecVirtualNetworkConfiguration)(ptr) = obj
		} else {
			*(*ClusterSpecVirtualNetworkConfiguration)(ptr) = ClusterSpecVirtualNetworkConfiguration{}
		}
	default:
		iter.ReportError("decode ClusterSpecVirtualNetworkConfiguration", "unexpected JSON type")
	}
}
