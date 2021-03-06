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
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettings{}).Type1()):                           GatewaySpecBgpSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettingsInstance0BGPPeeringAddress{}).Type1()): GatewaySpecBgpSettingsInstance0BGPPeeringAddressCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettingsInstance1BGPPeeringAddress{}).Type1()): GatewaySpecBgpSettingsInstance1BGPPeeringAddressCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewayConnectionSpecRouting{}).Type1()):                     GatewayConnectionSpecRoutingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewayConnectionSpecRoutingPropagatedRouteTable{}).Type1()): GatewayConnectionSpecRoutingPropagatedRouteTableCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecIpsecPolicy{}).Type1()):               ServerConfigurationSpecIpsecPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecRadius{}).Type1()):                    ServerConfigurationSpecRadiusCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecRadiusServer2{}).Type1()):             ServerConfigurationSpecRadiusServer2Codec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SiteSpecLinkBgp{}).Type1()):                                  SiteSpecLinkBgpCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettings{}).Type1()):                           GatewaySpecBgpSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettingsInstance0BGPPeeringAddress{}).Type1()): GatewaySpecBgpSettingsInstance0BGPPeeringAddressCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettingsInstance1BGPPeeringAddress{}).Type1()): GatewaySpecBgpSettingsInstance1BGPPeeringAddressCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewayConnectionSpecRouting{}).Type1()):                     GatewayConnectionSpecRoutingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewayConnectionSpecRoutingPropagatedRouteTable{}).Type1()): GatewayConnectionSpecRoutingPropagatedRouteTableCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecIpsecPolicy{}).Type1()):               ServerConfigurationSpecIpsecPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecRadius{}).Type1()):                    ServerConfigurationSpecRadiusCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecRadiusServer2{}).Type1()):             ServerConfigurationSpecRadiusServer2Codec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SiteSpecLinkBgp{}).Type1()):                                  SiteSpecLinkBgpCodec{},
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
type GatewaySpecBgpSettingsCodec struct {
}

func (GatewaySpecBgpSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecBgpSettings)(ptr) == nil
}

func (GatewaySpecBgpSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecBgpSettings)(ptr)
	var objs []GatewaySpecBgpSettings
	if obj != nil {
		objs = []GatewaySpecBgpSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecBgpSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecBgpSettings)(ptr) = GatewaySpecBgpSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecBgpSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecBgpSettings)(ptr) = objs[0]
			} else {
				*(*GatewaySpecBgpSettings)(ptr) = GatewaySpecBgpSettings{}
			}
		} else {
			*(*GatewaySpecBgpSettings)(ptr) = GatewaySpecBgpSettings{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecBgpSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecBgpSettings)(ptr) = obj
		} else {
			*(*GatewaySpecBgpSettings)(ptr) = GatewaySpecBgpSettings{}
		}
	default:
		iter.ReportError("decode GatewaySpecBgpSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySpecBgpSettingsInstance0BGPPeeringAddressCodec struct {
}

func (GatewaySpecBgpSettingsInstance0BGPPeeringAddressCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecBgpSettingsInstance0BGPPeeringAddress)(ptr) == nil
}

func (GatewaySpecBgpSettingsInstance0BGPPeeringAddressCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecBgpSettingsInstance0BGPPeeringAddress)(ptr)
	var objs []GatewaySpecBgpSettingsInstance0BGPPeeringAddress
	if obj != nil {
		objs = []GatewaySpecBgpSettingsInstance0BGPPeeringAddress{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettingsInstance0BGPPeeringAddress{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecBgpSettingsInstance0BGPPeeringAddressCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecBgpSettingsInstance0BGPPeeringAddress)(ptr) = GatewaySpecBgpSettingsInstance0BGPPeeringAddress{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecBgpSettingsInstance0BGPPeeringAddress

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettingsInstance0BGPPeeringAddress{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecBgpSettingsInstance0BGPPeeringAddress)(ptr) = objs[0]
			} else {
				*(*GatewaySpecBgpSettingsInstance0BGPPeeringAddress)(ptr) = GatewaySpecBgpSettingsInstance0BGPPeeringAddress{}
			}
		} else {
			*(*GatewaySpecBgpSettingsInstance0BGPPeeringAddress)(ptr) = GatewaySpecBgpSettingsInstance0BGPPeeringAddress{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecBgpSettingsInstance0BGPPeeringAddress

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettingsInstance0BGPPeeringAddress{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecBgpSettingsInstance0BGPPeeringAddress)(ptr) = obj
		} else {
			*(*GatewaySpecBgpSettingsInstance0BGPPeeringAddress)(ptr) = GatewaySpecBgpSettingsInstance0BGPPeeringAddress{}
		}
	default:
		iter.ReportError("decode GatewaySpecBgpSettingsInstance0BGPPeeringAddress", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySpecBgpSettingsInstance1BGPPeeringAddressCodec struct {
}

func (GatewaySpecBgpSettingsInstance1BGPPeeringAddressCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecBgpSettingsInstance1BGPPeeringAddress)(ptr) == nil
}

func (GatewaySpecBgpSettingsInstance1BGPPeeringAddressCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecBgpSettingsInstance1BGPPeeringAddress)(ptr)
	var objs []GatewaySpecBgpSettingsInstance1BGPPeeringAddress
	if obj != nil {
		objs = []GatewaySpecBgpSettingsInstance1BGPPeeringAddress{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettingsInstance1BGPPeeringAddress{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecBgpSettingsInstance1BGPPeeringAddressCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecBgpSettingsInstance1BGPPeeringAddress)(ptr) = GatewaySpecBgpSettingsInstance1BGPPeeringAddress{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecBgpSettingsInstance1BGPPeeringAddress

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettingsInstance1BGPPeeringAddress{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecBgpSettingsInstance1BGPPeeringAddress)(ptr) = objs[0]
			} else {
				*(*GatewaySpecBgpSettingsInstance1BGPPeeringAddress)(ptr) = GatewaySpecBgpSettingsInstance1BGPPeeringAddress{}
			}
		} else {
			*(*GatewaySpecBgpSettingsInstance1BGPPeeringAddress)(ptr) = GatewaySpecBgpSettingsInstance1BGPPeeringAddress{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecBgpSettingsInstance1BGPPeeringAddress

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBgpSettingsInstance1BGPPeeringAddress{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecBgpSettingsInstance1BGPPeeringAddress)(ptr) = obj
		} else {
			*(*GatewaySpecBgpSettingsInstance1BGPPeeringAddress)(ptr) = GatewaySpecBgpSettingsInstance1BGPPeeringAddress{}
		}
	default:
		iter.ReportError("decode GatewaySpecBgpSettingsInstance1BGPPeeringAddress", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewayConnectionSpecRoutingCodec struct {
}

func (GatewayConnectionSpecRoutingCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewayConnectionSpecRouting)(ptr) == nil
}

func (GatewayConnectionSpecRoutingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewayConnectionSpecRouting)(ptr)
	var objs []GatewayConnectionSpecRouting
	if obj != nil {
		objs = []GatewayConnectionSpecRouting{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewayConnectionSpecRouting{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewayConnectionSpecRoutingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewayConnectionSpecRouting)(ptr) = GatewayConnectionSpecRouting{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewayConnectionSpecRouting

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewayConnectionSpecRouting{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewayConnectionSpecRouting)(ptr) = objs[0]
			} else {
				*(*GatewayConnectionSpecRouting)(ptr) = GatewayConnectionSpecRouting{}
			}
		} else {
			*(*GatewayConnectionSpecRouting)(ptr) = GatewayConnectionSpecRouting{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewayConnectionSpecRouting

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewayConnectionSpecRouting{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewayConnectionSpecRouting)(ptr) = obj
		} else {
			*(*GatewayConnectionSpecRouting)(ptr) = GatewayConnectionSpecRouting{}
		}
	default:
		iter.ReportError("decode GatewayConnectionSpecRouting", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewayConnectionSpecRoutingPropagatedRouteTableCodec struct {
}

func (GatewayConnectionSpecRoutingPropagatedRouteTableCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewayConnectionSpecRoutingPropagatedRouteTable)(ptr) == nil
}

func (GatewayConnectionSpecRoutingPropagatedRouteTableCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewayConnectionSpecRoutingPropagatedRouteTable)(ptr)
	var objs []GatewayConnectionSpecRoutingPropagatedRouteTable
	if obj != nil {
		objs = []GatewayConnectionSpecRoutingPropagatedRouteTable{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewayConnectionSpecRoutingPropagatedRouteTable{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewayConnectionSpecRoutingPropagatedRouteTableCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewayConnectionSpecRoutingPropagatedRouteTable)(ptr) = GatewayConnectionSpecRoutingPropagatedRouteTable{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewayConnectionSpecRoutingPropagatedRouteTable

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewayConnectionSpecRoutingPropagatedRouteTable{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewayConnectionSpecRoutingPropagatedRouteTable)(ptr) = objs[0]
			} else {
				*(*GatewayConnectionSpecRoutingPropagatedRouteTable)(ptr) = GatewayConnectionSpecRoutingPropagatedRouteTable{}
			}
		} else {
			*(*GatewayConnectionSpecRoutingPropagatedRouteTable)(ptr) = GatewayConnectionSpecRoutingPropagatedRouteTable{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewayConnectionSpecRoutingPropagatedRouteTable

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewayConnectionSpecRoutingPropagatedRouteTable{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewayConnectionSpecRoutingPropagatedRouteTable)(ptr) = obj
		} else {
			*(*GatewayConnectionSpecRoutingPropagatedRouteTable)(ptr) = GatewayConnectionSpecRoutingPropagatedRouteTable{}
		}
	default:
		iter.ReportError("decode GatewayConnectionSpecRoutingPropagatedRouteTable", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServerConfigurationSpecIpsecPolicyCodec struct {
}

func (ServerConfigurationSpecIpsecPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServerConfigurationSpecIpsecPolicy)(ptr) == nil
}

func (ServerConfigurationSpecIpsecPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServerConfigurationSpecIpsecPolicy)(ptr)
	var objs []ServerConfigurationSpecIpsecPolicy
	if obj != nil {
		objs = []ServerConfigurationSpecIpsecPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecIpsecPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServerConfigurationSpecIpsecPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServerConfigurationSpecIpsecPolicy)(ptr) = ServerConfigurationSpecIpsecPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServerConfigurationSpecIpsecPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecIpsecPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServerConfigurationSpecIpsecPolicy)(ptr) = objs[0]
			} else {
				*(*ServerConfigurationSpecIpsecPolicy)(ptr) = ServerConfigurationSpecIpsecPolicy{}
			}
		} else {
			*(*ServerConfigurationSpecIpsecPolicy)(ptr) = ServerConfigurationSpecIpsecPolicy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ServerConfigurationSpecIpsecPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecIpsecPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ServerConfigurationSpecIpsecPolicy)(ptr) = obj
		} else {
			*(*ServerConfigurationSpecIpsecPolicy)(ptr) = ServerConfigurationSpecIpsecPolicy{}
		}
	default:
		iter.ReportError("decode ServerConfigurationSpecIpsecPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServerConfigurationSpecRadiusCodec struct {
}

func (ServerConfigurationSpecRadiusCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServerConfigurationSpecRadius)(ptr) == nil
}

func (ServerConfigurationSpecRadiusCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServerConfigurationSpecRadius)(ptr)
	var objs []ServerConfigurationSpecRadius
	if obj != nil {
		objs = []ServerConfigurationSpecRadius{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecRadius{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServerConfigurationSpecRadiusCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServerConfigurationSpecRadius)(ptr) = ServerConfigurationSpecRadius{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServerConfigurationSpecRadius

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecRadius{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServerConfigurationSpecRadius)(ptr) = objs[0]
			} else {
				*(*ServerConfigurationSpecRadius)(ptr) = ServerConfigurationSpecRadius{}
			}
		} else {
			*(*ServerConfigurationSpecRadius)(ptr) = ServerConfigurationSpecRadius{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ServerConfigurationSpecRadius

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecRadius{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ServerConfigurationSpecRadius)(ptr) = obj
		} else {
			*(*ServerConfigurationSpecRadius)(ptr) = ServerConfigurationSpecRadius{}
		}
	default:
		iter.ReportError("decode ServerConfigurationSpecRadius", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServerConfigurationSpecRadiusServer2Codec struct {
}

func (ServerConfigurationSpecRadiusServer2Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServerConfigurationSpecRadiusServer2)(ptr) == nil
}

func (ServerConfigurationSpecRadiusServer2Codec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServerConfigurationSpecRadiusServer2)(ptr)
	var objs []ServerConfigurationSpecRadiusServer2
	if obj != nil {
		objs = []ServerConfigurationSpecRadiusServer2{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecRadiusServer2{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServerConfigurationSpecRadiusServer2Codec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServerConfigurationSpecRadiusServer2)(ptr) = ServerConfigurationSpecRadiusServer2{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServerConfigurationSpecRadiusServer2

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecRadiusServer2{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServerConfigurationSpecRadiusServer2)(ptr) = objs[0]
			} else {
				*(*ServerConfigurationSpecRadiusServer2)(ptr) = ServerConfigurationSpecRadiusServer2{}
			}
		} else {
			*(*ServerConfigurationSpecRadiusServer2)(ptr) = ServerConfigurationSpecRadiusServer2{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ServerConfigurationSpecRadiusServer2

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServerConfigurationSpecRadiusServer2{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ServerConfigurationSpecRadiusServer2)(ptr) = obj
		} else {
			*(*ServerConfigurationSpecRadiusServer2)(ptr) = ServerConfigurationSpecRadiusServer2{}
		}
	default:
		iter.ReportError("decode ServerConfigurationSpecRadiusServer2", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SiteSpecLinkBgpCodec struct {
}

func (SiteSpecLinkBgpCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SiteSpecLinkBgp)(ptr) == nil
}

func (SiteSpecLinkBgpCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SiteSpecLinkBgp)(ptr)
	var objs []SiteSpecLinkBgp
	if obj != nil {
		objs = []SiteSpecLinkBgp{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SiteSpecLinkBgp{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SiteSpecLinkBgpCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SiteSpecLinkBgp)(ptr) = SiteSpecLinkBgp{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SiteSpecLinkBgp

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SiteSpecLinkBgp{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SiteSpecLinkBgp)(ptr) = objs[0]
			} else {
				*(*SiteSpecLinkBgp)(ptr) = SiteSpecLinkBgp{}
			}
		} else {
			*(*SiteSpecLinkBgp)(ptr) = SiteSpecLinkBgp{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj SiteSpecLinkBgp

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SiteSpecLinkBgp{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*SiteSpecLinkBgp)(ptr) = obj
		} else {
			*(*SiteSpecLinkBgp)(ptr) = SiteSpecLinkBgp{}
		}
	default:
		iter.ReportError("decode SiteSpecLinkBgp", "unexpected JSON type")
	}
}
