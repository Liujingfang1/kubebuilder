package generators
package generators

import (
"bytes"
"fmt"
"io"
"path/filepath"
"reflect"
"sort"
"strings"

"k8s.io/gengo/args"
"k8s.io/gengo/generator"
"k8s.io/gengo/namer"
"k8s.io/gengo/types"
openapi "k8s.io/kube-openapi/pkg/common"
"github.com/go-openapi/spec"

"github.com/golang/glog"
)


func TypeToSchemaProps(types.Type) (spec.SchemaProps, string) {
	
} 