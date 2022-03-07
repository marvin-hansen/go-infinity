#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

source "../../scripts/bash_shared/convert_functions.sh"

echo "Hello"

[ $# -eq 0 ] && {
  echo "Usage: package (lowercase) Entity (CamelCase):"; exit 1;
}

namevar=$1
name=$2

fileExt=".go"
snakeName=$(convert_camel_to_snake_case_func "$name") # snake_case

fileName=$(to_lowercase "$snakeName")"$fileExt"

pack_name=$namevar

func_name="NewRequestFor"$name"ApiVersion"

cat << EOF > $fileName
package $pack_name

import (
	"fmt"
	"go-infinity/rest/flags"
	"net/http"
)

const defaultVersion = flags.VERSION

//**// Request //**//

func NewRequestFor$name() *RequestFor$name {
	return &RequestFor$name{defaultVersion}
}

func $func_name(apiVersion string) *RequestFor$name {
	return &RequestFor$name{apiVersion}
}

type RequestFor$name struct {
	version string
}

func (req *RequestFor$name) Path() string {
  // @FIXME
	return "/" + req.version + "/"
}

func (req *RequestFor$name) Method() string {
	return http.MethodGet
}

func (req *RequestFor$name) Query() string {
	return ""
}

func (req *RequestFor$name) Payload() []byte {
	return nil
}

//**// Response //**//

func NewResponseFor$name() *ResponseFor$name {
	return new(ResponseFor$name)
}

func (r ResponseFor$name) String() string {
  // @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}

type ResponseFor$name struct {
  // @FIXME
	Field string `json:"field"`
}


EOF

## Generate build files & update bazel dependencies
command bazel run //:gazelle
## Build all sources
command bazel build //:build