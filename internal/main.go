package internal

import (
	"log"
	"strings"

	"github.com/ZoeySimone/sidra/options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GenerateFile(gen *protogen.Plugin, file *protogen.File, provider *Provider) {
	parseFileOptions(file, provider)

	services := []Service{}
	for _, service := range file.Services {
		services = append(services, genService(service))
	}
	provider.Services = append(provider.Services, services...)

	res := &Resource{}
	if len(services) > 0 {
		res.HasServices = true
	}

	res.PackageName = string(file.GoPackageName)
	provider.PackageName = string(file.GoPackageName)

	for _, msg := range file.Messages {
		message := genMessage(msg)
		if message != nil {
			res.Messages = append(res.Messages, message)
		}
	}

	parseCRUD(res, provider.Services)

	// generate file if it contains services or messages
	if len(services) > 0 || len(res.Messages) > 0 {
		filename := file.GeneratedFilenamePrefix + "_resource.go"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		WriteTemplate("templates/resource.tmpl", g, res)
	}

	// add generated resources to the provider
	provider.Resources = append(provider.Resources, *res)
}

func parseFileOptions(f *protogen.File, provider *Provider) {
	opts, ok := proto.GetExtension(f.Proto.Options, options.E_Provider).(*options.FileOptions)
	if !ok {
		return
	}

	if provider.ProviderName == "" {
		name := opts.GetName()
		if name != "" {
			provider.ProviderName = name
		} else {
			log.Print("No provider name specified")
		}
	}
}

func parseMethodOptions(m *Method, o protoreflect.ProtoMessage) {
	opts, ok := proto.GetExtension(o.(*descriptorpb.MethodOptions), options.E_Method).(*options.MethodOptions)
	if !ok {
		return
	}

	if opts != nil {
		m.Operation = opts.GetOperation()
		m.Resource = opts.GetResource()
	}
}

func genService(s *protogen.Service) Service {
	service := Service{}
	service.ServiceName = s.GoName
	for _, method := range s.Methods {
		m := &Method{}
		m.ServiceName = service.ServiceName
		m.Function = method.GoName
		m.ParameterType = method.Input.GoIdent.GoName
		m.ReturnType = method.Output.GoIdent.GoName
		parseMethodOptions(m, method.Desc.Options())

		service.Methods = append(service.Methods, *m)
	}

	return service
}

func parseMessageOptions(m *Message, o protoreflect.ProtoMessage) {
	opts, ok := proto.GetExtension(o.(*descriptorpb.MessageOptions), options.E_Resource).(*options.MessageOptions)
	if !ok {
		return
	}

	if opts != nil {
		m.ignore = opts.GetIgnore()
	}
}

func renameField(name string) string {
	return strings.ToLower(name)
}

func parseFieldOptions(f *MessageField, p protoreflect.ProtoMessage) {
	opts, ok := p.(*descriptorpb.FieldOptions)
	if !ok {
		return
	}
	if opts != nil {
		deprecated := opts.Deprecated
		if deprecated != nil {
			f.Deprecated = *deprecated
		}

		extOpts, ok := proto.GetExtension(opts, options.E_Field).(*options.FieldOptions)
		if !ok {
			return
		}

		if extOpts != nil {
			f.Sensitive = extOpts.GetSensitive()
			f.ID = extOpts.GetID()
			behaviour := strings.ToLower(extOpts.GetBehaviour())

			switch behaviour {
			case "computed":
				f.Computed = true
			case "optional":
				f.Optional = true
			case "required":
				f.Required = true
			}
		}
	}
}

func genMessage(m *protogen.Message) *Message {
	msg := &Message{}
	parseMessageOptions(msg, m.Desc.Options())
	if msg.ignore || len(m.Fields) == 0 {
		return nil
	}

	msg.GoName = m.GoIdent.GoName
	msg.Name = renameField(msg.GoName)
	for _, field := range m.Fields {

		f := &MessageField{}
		f.Name = renameField(field.GoName)
		f.GoName = field.GoName
		f.Description = strings.TrimSpace(strings.TrimPrefix(field.Comments.Trailing.String(), "//"))

		parseFieldOptions(f, field.Desc.Options())
		if !f.Computed && !f.Required {
			// if behaviour not set manually use optional keyword
			f.Optional = field.Desc.HasOptionalKeyword()
		}
		setRequired(f)

		parseType(f, field)

		if f.ID {
			msg.IDField = f
		}
		msg.Fields = append(msg.Fields, *f)
	}

	return msg
}

func parseCRUD(r *Resource, s []Service) {
	for i := range r.Messages {
		findCRUDFunc(r.Messages[i], s)
	}
}

// find appropriate methods for the resource
func findCRUDFunc(m *Message, services []Service) {
	for _, s := range services {
		for i, method := range s.Methods {
			if method.Resource == m.GoName {
				switch method.Operation {
				case "Create":
					m.Create = &s.Methods[i]
				case "Read":
					m.Read = &s.Methods[i]
				case "Update":
					m.Update = &s.Methods[i]
				case "Delete":
					m.Delete = &s.Methods[i]
				case "Exists":
					m.Exists = &s.Methods[i]
				}
			}
		}
	}
}

func setRequired(f *MessageField) {
	if !f.Optional && !f.Computed {
		// assume any field which is not optional or computed is required
		f.Required = true
	}
}

func parseType(f *MessageField, field *protogen.Field) {
	if field.Desc.IsList() {
		f.Type = "schema.TypeList"
		f.Elem = &Schema{
			Type: parseKind(field.Desc.Kind()),
		}
		f.GoType = "[]" + getGoType(field.Desc.Kind())
	} else if field.Desc.IsMap() {
		f.Type = "schema.TypeMap"
		f.Elem = &Schema{
			Type: parseKind(field.Desc.MapValue().Kind()),
		}
		f.GoType = "map[" + getGoType(field.Desc.MapKey().Kind()) + "]" + getGoType(field.Desc.MapValue().Kind())
	} else if field.Desc.Kind() == protoreflect.MessageKind {
		f.Type = "schema.TypeMap"
		f.SubSchema = genMessage(field.Message)
	} else {
		f.Type = parseKind(field.Desc.Kind())

		// add a pointer to optional field types
		if f.Optional {
			f.GoType = "*"
		}

		f.GoType += getGoType(field.Desc.Kind())
	}
}

func parseKind(kind protoreflect.Kind) string {
	switch kind {
	case protoreflect.Int32Kind, protoreflect.Int64Kind:
		return "schema.TypeInt"
	case protoreflect.Uint32Kind, protoreflect.Sfixed64Kind:
		return "schema.TypeInt"
	case protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind, protoreflect.Sint64Kind:
		return "schema.TypeInt"
	case protoreflect.StringKind:
		return "schema.TypeString"
	case protoreflect.BoolKind:
		return "schema.TypeBool"
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return "schema.TypeFloat"
	case protoreflect.EnumKind:
		return "schema.TypeInt"
	}
	return "schema.TypeInvalid"
}

func getGoType(kind protoreflect.Kind) string {
	switch kind {
	case protoreflect.BoolKind:
		return "bool"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "uint64"
	case protoreflect.FloatKind:
		return "float32"
	case protoreflect.DoubleKind:
		return "float64"
	case protoreflect.StringKind:
		return "string"
	case protoreflect.BytesKind:
		return "[]byte"
	case protoreflect.EnumKind:
		return "int64"
	}
	return "UnknownType"
}
