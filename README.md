# Sidra - Protobuf to Terraform provider
[![Sidra Test and Lint](https://github.com/ZoeySimone/sidra/actions/workflows/workflow.yml/badge.svg?branch=main)](https://github.com/ZoeySimone/sidra/actions/workflows/workflow.yml)
## Install

Most recent release for linux/amd64 can be downloaded from [https://github.com/ZoeySimone/sidra/releases]() and placed in the PATH.

Alternatively to build code

    git clone https://github.com/ZoeySimone/sidra
    cd sidra
    make sidra


## Usage

    protoc ./example.proto --go_out=. --sidra_out=.

### Protobuf Options

To improve the chances of a complete Terraform provider generation and improved readability of the resulting code it is possible to use Protobuf Options to manually specify certain information. To use the options [sidra.proto](./options/sidra.proto) file must be imported into the protobuf file.

#### FileOptions

Terraform provider name is a required option to generate a Terraform plugin successfully
    option (sidra.Provider = {Name: "provider_name"}) 
#### MessageOptions

Specify a custom name for the generated resource:

    option (sidra.Resource) = {ResourceName: "User"};

Prevent message from generating a resource

    option (sidra.Resource) = {Ignore: true};

#### FieldOptions

Mark field as an ID, option is necessary if there is no field named "id" in the message

    option (sidra.Field) = {ID: true};

Prevent field from being added into the schema

    option (sidra.Field) = {Ignore: true};

Mark field as sensitive

    option (sidra.Field) = {Sensitive: true};

Define whether the value is Computed, Required or Optional
    options (sidra.Field) = {Behaviour: "computed"};

#### MethodOptions

Define type of CRUD(E) operation for the Method

    option (sidra.Method) = {Operation: "Create" Resource: "User"};

## Templates

Default templates are embedded into the binary, to make changes to the templates it is necessary to change the files in [internal/templates](./internal/templates) and build the binary.