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
		jsoniter.MustGetKind(reflect2.TypeOf(NamespaceSpecIdentity{}).Type1()):                 NamespaceSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SubscriptionRuleSpecCorrelationFilter{}).Type1()): SubscriptionRuleSpecCorrelationFilterCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(NamespaceSpecIdentity{}).Type1()):                 NamespaceSpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SubscriptionRuleSpecCorrelationFilter{}).Type1()): SubscriptionRuleSpecCorrelationFilterCodec{},
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
type NamespaceSpecIdentityCodec struct {
}

func (NamespaceSpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*NamespaceSpecIdentity)(ptr) == nil
}

func (NamespaceSpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*NamespaceSpecIdentity)(ptr)
	var objs []NamespaceSpecIdentity
	if obj != nil {
		objs = []NamespaceSpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NamespaceSpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (NamespaceSpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*NamespaceSpecIdentity)(ptr) = NamespaceSpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []NamespaceSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NamespaceSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*NamespaceSpecIdentity)(ptr) = objs[0]
			} else {
				*(*NamespaceSpecIdentity)(ptr) = NamespaceSpecIdentity{}
			}
		} else {
			*(*NamespaceSpecIdentity)(ptr) = NamespaceSpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj NamespaceSpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NamespaceSpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*NamespaceSpecIdentity)(ptr) = obj
		} else {
			*(*NamespaceSpecIdentity)(ptr) = NamespaceSpecIdentity{}
		}
	default:
		iter.ReportError("decode NamespaceSpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SubscriptionRuleSpecCorrelationFilterCodec struct {
}

func (SubscriptionRuleSpecCorrelationFilterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SubscriptionRuleSpecCorrelationFilter)(ptr) == nil
}

func (SubscriptionRuleSpecCorrelationFilterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SubscriptionRuleSpecCorrelationFilter)(ptr)
	var objs []SubscriptionRuleSpecCorrelationFilter
	if obj != nil {
		objs = []SubscriptionRuleSpecCorrelationFilter{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SubscriptionRuleSpecCorrelationFilter{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SubscriptionRuleSpecCorrelationFilterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SubscriptionRuleSpecCorrelationFilter)(ptr) = SubscriptionRuleSpecCorrelationFilter{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SubscriptionRuleSpecCorrelationFilter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SubscriptionRuleSpecCorrelationFilter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SubscriptionRuleSpecCorrelationFilter)(ptr) = objs[0]
			} else {
				*(*SubscriptionRuleSpecCorrelationFilter)(ptr) = SubscriptionRuleSpecCorrelationFilter{}
			}
		} else {
			*(*SubscriptionRuleSpecCorrelationFilter)(ptr) = SubscriptionRuleSpecCorrelationFilter{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj SubscriptionRuleSpecCorrelationFilter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SubscriptionRuleSpecCorrelationFilter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*SubscriptionRuleSpecCorrelationFilter)(ptr) = obj
		} else {
			*(*SubscriptionRuleSpecCorrelationFilter)(ptr) = SubscriptionRuleSpecCorrelationFilter{}
		}
	default:
		iter.ReportError("decode SubscriptionRuleSpecCorrelationFilter", "unexpected JSON type")
	}
}
