package internal

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func Test_parseKind(t *testing.T) {
	tests := []struct {
		name string
		kind protoreflect.Kind
		want string
	}{
		{name: "bool", kind: protoreflect.BoolKind, want: "schema.TypeBool"},
		{name: "int32", kind: protoreflect.Int32Kind, want: "schema.TypeInt"},
		{name: "sint32", kind: protoreflect.Sint32Kind, want: "schema.TypeInt"},
		{name: "uint32", kind: protoreflect.Uint32Kind, want: "schema.TypeInt"},
		{name: "int64", kind: protoreflect.Int64Kind, want: "schema.TypeInt"},
		{name: "sint64", kind: protoreflect.Sint64Kind, want: "schema.TypeInt"},
		{name: "uint64", kind: protoreflect.Uint64Kind, want: "schema.TypeInvalid"},
		{name: "sfixed32", kind: protoreflect.Sfixed32Kind, want: "schema.TypeInt"},
		{name: "fixed32", kind: protoreflect.Fixed32Kind, want: "schema.TypeInt"},
		{name: "sfixed64", kind: protoreflect.Sfixed64Kind, want: "schema.TypeInt"},
		{name: "fixed64", kind: protoreflect.Fixed64Kind, want: "schema.TypeInvalid"},
		{name: "float", kind: protoreflect.FloatKind, want: "schema.TypeFloat"},
		{name: "double", kind: protoreflect.DoubleKind, want: "schema.TypeFloat"},
		{name: "string", kind: protoreflect.StringKind, want: "schema.TypeString"},
		{name: "message", kind: protoreflect.MessageKind, want: "schema.TypeInvalid"},
		{name: "enum", kind: protoreflect.EnumKind, want: "schema.TypeInt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseKind(tt.kind); got != tt.want {
				t.Errorf("parseType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGoType(t *testing.T) {
	tests := []struct {
		name string
		kind protoreflect.Kind
		want string
	}{
		{name: "bool", kind: protoreflect.BoolKind, want: "bool"},
		{name: "int32", kind: protoreflect.Int32Kind, want: "int32"},
		{name: "sint32", kind: protoreflect.Sint32Kind, want: "int32"},
		{name: "uint32", kind: protoreflect.Uint32Kind, want: "uint32"},
		{name: "int64", kind: protoreflect.Int64Kind, want: "int64"},
		{name: "sint64", kind: protoreflect.Sint64Kind, want: "int64"},
		{name: "uint64", kind: protoreflect.Uint64Kind, want: "uint64"},
		{name: "sfixed32", kind: protoreflect.Sfixed32Kind, want: "int32"},
		{name: "fixed32", kind: protoreflect.Fixed32Kind, want: "uint32"},
		{name: "sfixed64", kind: protoreflect.Sfixed64Kind, want: "int64"},
		{name: "fixed64", kind: protoreflect.Fixed64Kind, want: "uint64"},
		{name: "float", kind: protoreflect.FloatKind, want: "float32"},
		{name: "double", kind: protoreflect.DoubleKind, want: "float64"},
		{name: "string", kind: protoreflect.StringKind, want: "string"},
		{name: "message", kind: protoreflect.MessageKind, want: "UnknownType"},
		{name: "enum", kind: protoreflect.EnumKind, want: "int64"},
		{name: "bytes", kind: protoreflect.BytesKind, want: "[]byte"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGoType(tt.kind); got != tt.want {
				t.Errorf("getGoType() = %v, want %v", got, tt.want)
			}
		})
	}
}
