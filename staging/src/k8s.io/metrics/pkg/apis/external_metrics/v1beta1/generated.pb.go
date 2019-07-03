/*
Copyright The Kubernetes Authors.

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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/external_metrics/v1beta1/generated.proto

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/external_metrics/v1beta1/generated.proto

	It has these top-level messages:
		ExternalMetricValue
		ExternalMetricValueList
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import sortkeys "github.com/gogo/protobuf/sortkeys"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *ExternalMetricValue) Reset()                    { *m = ExternalMetricValue{} }
func (*ExternalMetricValue) ProtoMessage()               {}
func (*ExternalMetricValue) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *ExternalMetricValueList) Reset()                    { *m = ExternalMetricValueList{} }
func (*ExternalMetricValueList) ProtoMessage()               {}
func (*ExternalMetricValueList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func init() {
	proto.RegisterType((*ExternalMetricValue)(nil), "k8s.io.metrics.pkg.apis.external_metrics.v1beta1.ExternalMetricValue")
	proto.RegisterType((*ExternalMetricValueList)(nil), "k8s.io.metrics.pkg.apis.external_metrics.v1beta1.ExternalMetricValueList")
}
func (m *ExternalMetricValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalMetricValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.MetricName)))
	i += copy(dAtA[i:], m.MetricName)
	if len(m.MetricLabels) > 0 {
		keysForMetricLabels := make([]string, 0, len(m.MetricLabels))
		for k := range m.MetricLabels {
			keysForMetricLabels = append(keysForMetricLabels, string(k))
		}
		sortkeys.Strings(keysForMetricLabels)
		for _, k := range keysForMetricLabels {
			dAtA[i] = 0x12
			i++
			v := m.MetricLabels[string(k)]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Timestamp.Size()))
	n1, err := m.Timestamp.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.WindowSeconds != nil {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.WindowSeconds))
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Value.Size()))
	n2, err := m.Value.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *ExternalMetricValueList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalMetricValueList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n3, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ExternalMetricValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.MetricName)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.MetricLabels) > 0 {
		for k, v := range m.MetricLabels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	l = m.Timestamp.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.WindowSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.WindowSeconds))
	}
	l = m.Value.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ExternalMetricValueList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ExternalMetricValue) String() string {
	if this == nil {
		return "nil"
	}
	keysForMetricLabels := make([]string, 0, len(this.MetricLabels))
	for k := range this.MetricLabels {
		keysForMetricLabels = append(keysForMetricLabels, k)
	}
	sortkeys.Strings(keysForMetricLabels)
	mapStringForMetricLabels := "map[string]string{"
	for _, k := range keysForMetricLabels {
		mapStringForMetricLabels += fmt.Sprintf("%v: %v,", k, this.MetricLabels[k])
	}
	mapStringForMetricLabels += "}"
	s := strings.Join([]string{`&ExternalMetricValue{`,
		`MetricName:` + fmt.Sprintf("%v", this.MetricName) + `,`,
		`MetricLabels:` + mapStringForMetricLabels + `,`,
		`Timestamp:` + strings.Replace(strings.Replace(this.Timestamp.String(), "Time", "k8s_io_apimachinery_pkg_apis_meta_v1.Time", 1), `&`, ``, 1) + `,`,
		`WindowSeconds:` + valueToStringGenerated(this.WindowSeconds) + `,`,
		`Value:` + strings.Replace(strings.Replace(this.Value.String(), "Quantity", "k8s_io_apimachinery_pkg_api_resource.Quantity", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ExternalMetricValueList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ExternalMetricValueList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "ExternalMetricValue", "ExternalMetricValue", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ExternalMetricValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExternalMetricValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExternalMetricValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetricName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricLabels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetricLabels == nil {
				m.MetricLabels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MetricLabels[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowSeconds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WindowSeconds = &v
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExternalMetricValueList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExternalMetricValueList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExternalMetricValueList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, ExternalMetricValue{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/external_metrics/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0x89, 0xd1, 0x66, 0xda, 0x42, 0xb3, 0x16, 0x0c, 0x39, 0x6c, 0x42, 0x4f, 0x51,
	0xe8, 0x8c, 0x09, 0x22, 0xc5, 0x8b, 0xb0, 0x90, 0x83, 0xd0, 0x08, 0x6e, 0x8b, 0xc5, 0x1f, 0x20,
	0x93, 0xcd, 0x73, 0x33, 0x26, 0xfb, 0x83, 0x99, 0xd9, 0xd4, 0xdc, 0xfc, 0x13, 0xf4, 0xbf, 0xca,
	0xb1, 0xc7, 0x9e, 0x82, 0xd9, 0xfe, 0x19, 0x5e, 0x64, 0x77, 0x67, 0x9b, 0x98, 0xd4, 0x16, 0xc1,
	0xdb, 0xbc, 0x79, 0xf3, 0xbe, 0xef, 0xf3, 0xbe, 0xf3, 0xf0, 0xbb, 0xd1, 0x91, 0x24, 0x3c, 0xa0,
	0xa3, 0xa8, 0x0f, 0xc2, 0x07, 0x05, 0x92, 0x4e, 0xc0, 0x1f, 0x04, 0x82, 0xea, 0x84, 0x07, 0x4a,
	0x70, 0x47, 0xd2, 0x70, 0xe4, 0x52, 0x16, 0x72, 0x49, 0xe1, 0xab, 0x02, 0xe1, 0xb3, 0xf1, 0xa7,
	0x3c, 0x33, 0x69, 0xf7, 0x41, 0xb1, 0x36, 0x75, 0xc1, 0x07, 0xc1, 0x14, 0x0c, 0x48, 0x28, 0x02,
	0x15, 0x18, 0x4f, 0x33, 0x05, 0xa2, 0xdf, 0x91, 0x70, 0xe4, 0x92, 0x44, 0x81, 0xac, 0x2b, 0x10,
	0xad, 0x50, 0x3f, 0x74, 0xb9, 0x1a, 0x46, 0x7d, 0xe2, 0x04, 0x1e, 0x75, 0x03, 0x37, 0xa0, 0xa9,
	0x50, 0x3f, 0xfa, 0x9c, 0x46, 0x69, 0x90, 0x9e, 0xb2, 0x06, 0xf5, 0x67, 0x1a, 0x91, 0x85, 0xdc,
	0x63, 0xce, 0x90, 0xfb, 0x20, 0xa6, 0x39, 0x27, 0x15, 0x20, 0x83, 0x48, 0x38, 0xb0, 0x8e, 0x75,
	0x6b, 0x95, 0x4c, 0xc6, 0x65, 0x74, 0xb2, 0x31, 0x4c, 0x9d, 0xfe, 0xad, 0x4a, 0x44, 0xbe, 0xe2,
	0xde, 0x66, 0x9b, 0xe7, 0x77, 0x15, 0x48, 0x67, 0x08, 0x1e, 0x5b, 0xaf, 0x3b, 0xf8, 0x55, 0xc2,
	0x0f, 0xbb, 0xda, 0xa0, 0x5e, 0xea, 0xcf, 0x5b, 0x36, 0x8e, 0xc0, 0xe8, 0x60, 0x9c, 0xd9, 0xf5,
	0x9a, 0x79, 0x50, 0x43, 0x4d, 0xd4, 0xaa, 0x58, 0xc6, 0x6c, 0xde, 0x28, 0xc4, 0xf3, 0x06, 0xee,
	0x5d, 0x67, 0xec, 0x95, 0x57, 0xc6, 0x0f, 0x84, 0x77, 0xb2, 0xf0, 0x98, 0xf5, 0x61, 0x2c, 0x6b,
	0xc5, 0x66, 0xa9, 0xb5, 0xdd, 0x39, 0x23, 0xff, 0xfa, 0x33, 0xe4, 0x06, 0x22, 0xd2, 0x5b, 0x51,
	0xee, 0xfa, 0x4a, 0x4c, 0xad, 0x7d, 0xcd, 0xb3, 0xb3, 0x9a, 0xb2, 0xff, 0x40, 0x30, 0x3e, 0xe0,
	0x4a, 0x32, 0xbe, 0x54, 0xcc, 0x0b, 0x6b, 0xa5, 0x26, 0x6a, 0x6d, 0x77, 0x9e, 0xe4, 0x3c, 0xab,
	0x5e, 0x2d, 0xa1, 0x92, 0x2f, 0x21, 0x93, 0x36, 0x39, 0xe5, 0x1e, 0x58, 0x55, 0xdd, 0xa2, 0x72,
	0x9a, 0x8b, 0xd8, 0x4b, 0x3d, 0xe3, 0x31, 0xbe, 0x7f, 0xce, 0xfd, 0x41, 0x70, 0x5e, 0xbb, 0xd7,
	0x44, 0xad, 0x92, 0x55, 0x8d, 0xe7, 0x8d, 0xdd, 0xb3, 0xf4, 0xe6, 0x04, 0x9c, 0xc0, 0x1f, 0x48,
	0x5b, 0x3f, 0x30, 0x4e, 0x70, 0x79, 0x92, 0x8c, 0x51, 0x2b, 0xa7, 0x0c, 0xe4, 0x36, 0x06, 0x92,
	0x2f, 0x13, 0x79, 0x13, 0x31, 0x5f, 0x71, 0x35, 0xb5, 0x76, 0x35, 0x47, 0x39, 0xf5, 0xc2, 0xce,
	0xb4, 0xea, 0x2f, 0x71, 0x75, 0xc3, 0x15, 0x63, 0x0f, 0x97, 0x46, 0x30, 0xcd, 0xbe, 0xcc, 0x4e,
	0x8e, 0xc6, 0x7e, 0xde, 0xbb, 0x98, 0xde, 0x65, 0xc1, 0x8b, 0xe2, 0x11, 0x3a, 0xb8, 0x42, 0xf8,
	0xd1, 0x0d, 0x5e, 0x1f, 0x73, 0xa9, 0x8c, 0x8f, 0x78, 0x2b, 0xb1, 0x62, 0xc0, 0x14, 0x4b, 0xc5,
	0xee, 0x80, 0x5e, 0x1a, 0x97, 0x54, 0xf7, 0x40, 0x31, 0x6b, 0x4f, 0x43, 0x6f, 0xe5, 0x37, 0xf6,
	0xb5, 0xa2, 0xf1, 0x05, 0x97, 0xb9, 0x02, 0x2f, 0xdf, 0x91, 0xee, 0x7f, 0xd9, 0x91, 0xa5, 0x4d,
	0xaf, 0x12, 0x6d, 0x3b, 0x6b, 0x61, 0x1d, 0xce, 0x16, 0x66, 0xe1, 0x62, 0x61, 0x16, 0x2e, 0x17,
	0x66, 0xe1, 0x5b, 0x6c, 0xa2, 0x59, 0x6c, 0xa2, 0x8b, 0xd8, 0x44, 0x97, 0xb1, 0x89, 0x7e, 0xc6,
	0x26, 0xfa, 0x7e, 0x65, 0x16, 0xde, 0x3f, 0xd0, 0xc2, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbf,
	0x2a, 0x7e, 0x2e, 0xa4, 0x04, 0x00, 0x00,
}
