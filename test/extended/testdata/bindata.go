// Code generated by go-bindata.
// sources:
// test/extended/testdata/apibinding/api_binding.yaml
// test/extended/testdata/apibinding/api_export.yaml
// test/extended/testdata/apibinding/api_rs.yaml
// test/extended/testdata/apibinding/cowboy.yaml
// test/extended/testdata/apibinding/role_hack.yaml
// test/extended/testdata/quota/quota.yaml
// test/extended/testdata/workspace/api_rs.yaml
// DO NOT EDIT!

package testdata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _testExtendedTestdataApibindingApi_bindingYaml = []byte(`apiVersion: apis.kcp.dev/v1alpha1
kind: APIBinding
metadata:
  name: cowboys
spec:
  reference:
    workspace:
      exportName: today-cowboys
`)

func testExtendedTestdataApibindingApi_bindingYamlBytes() ([]byte, error) {
	return _testExtendedTestdataApibindingApi_bindingYaml, nil
}

func testExtendedTestdataApibindingApi_bindingYaml() (*asset, error) {
	bytes, err := testExtendedTestdataApibindingApi_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/extended/testdata/apibinding/api_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testExtendedTestdataApibindingApi_exportYaml = []byte(`apiVersion: apis.kcp.dev/v1alpha1
kind: APIExport
metadata:
  name: today-cowboys
spec:
  latestResourceSchemas: 
    - today.cowboys.wildwest.dev`)

func testExtendedTestdataApibindingApi_exportYamlBytes() ([]byte, error) {
	return _testExtendedTestdataApibindingApi_exportYaml, nil
}

func testExtendedTestdataApibindingApi_exportYaml() (*asset, error) {
	bytes, err := testExtendedTestdataApibindingApi_exportYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/extended/testdata/apibinding/api_export.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testExtendedTestdataApibindingApi_rsYaml = []byte(`apiVersion: apis.kcp.dev/v1alpha1
kind: APIResourceSchema
metadata:
  name: today.cowboys.wildwest.dev
spec:
  group: wildwest.dev
  names:
    kind: Cowboy
    listKind: CowboyList
    plural: cowboys
    singular: cowboy
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      description: Cowboy is part of the wild west
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: CowboySpec holds the desired state of the Cowboy.
          properties:
            intent:
              type: string
          type: object
        status:
          description: CowboyStatus communicates the observed state of the Cowboy.
          properties:
            result:
              type: string
          type: object
      type: object
    served: true
    storage: true
    subresources:
      status: {}`)

func testExtendedTestdataApibindingApi_rsYamlBytes() ([]byte, error) {
	return _testExtendedTestdataApibindingApi_rsYaml, nil
}

func testExtendedTestdataApibindingApi_rsYaml() (*asset, error) {
	bytes, err := testExtendedTestdataApibindingApi_rsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/extended/testdata/apibinding/api_rs.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testExtendedTestdataApibindingCowboyYaml = []byte(`apiVersion: wildwest.dev/v1alpha1
kind: Cowboy
metadata:
  name: one
spec:
  intent: one
`)

func testExtendedTestdataApibindingCowboyYamlBytes() ([]byte, error) {
	return _testExtendedTestdataApibindingCowboyYaml, nil
}

func testExtendedTestdataApibindingCowboyYaml() (*asset, error) {
	bytes, err := testExtendedTestdataApibindingCowboyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/extended/testdata/apibinding/cowboy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testExtendedTestdataApibindingRole_hackYaml = []byte(`---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: bind-apiexport-spi
rules:
- apiGroups:
  - apis.kcp.dev
  resourceNames:
  - today-cowboys
  resources:
  - apiexports
  verbs:
  - bind
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: bind-apiexport-spi
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: bind-apiexport-spi
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:authenticated`)

func testExtendedTestdataApibindingRole_hackYamlBytes() ([]byte, error) {
	return _testExtendedTestdataApibindingRole_hackYaml, nil
}

func testExtendedTestdataApibindingRole_hackYaml() (*asset, error) {
	bytes, err := testExtendedTestdataApibindingRole_hackYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/extended/testdata/apibinding/role_hack.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testExtendedTestdataQuotaQuotaYaml = []byte(`apiVersion: v1
kind: ResourceQuota
metadata:
  annotations:
    experimental.quota.kcp.dev/cluster-scoped: "true"
  name: quota
spec:
  hard:
    count/configmaps: "$(NUM_CONFIGMAP)"
    count/secrets: "$(NUM_SECRET)"`)

func testExtendedTestdataQuotaQuotaYamlBytes() ([]byte, error) {
	return _testExtendedTestdataQuotaQuotaYaml, nil
}

func testExtendedTestdataQuotaQuotaYaml() (*asset, error) {
	bytes, err := testExtendedTestdataQuotaQuotaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/extended/testdata/quota/quota.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testExtendedTestdataWorkspaceApi_rsYaml = []byte(`apiVersion: apis.kcp.dev/v1alpha1
kind: APIResourceSchema
metadata:
  name: v1.widgets.example.io
spec:
  group: example.io
  names:
    kind: Widget
    listKind: WidgetList
    plural: widgets
    singular: widget
  scope: Cluster
  versions:
  - name: v1
    schema:
      properties:
        apiVersion:
          type: string
        kind:
          type: string
        metadata:
          type: object
        spec:
          properties:
            firstName:
              type: string
            lastName:
              type: string
          type: object
        status:
          properties:
            phase:
              type: string
          type: object
      type: object
    served: true
    storage: true
    subresources:
      status: {}
`)

func testExtendedTestdataWorkspaceApi_rsYamlBytes() ([]byte, error) {
	return _testExtendedTestdataWorkspaceApi_rsYaml, nil
}

func testExtendedTestdataWorkspaceApi_rsYaml() (*asset, error) {
	bytes, err := testExtendedTestdataWorkspaceApi_rsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/extended/testdata/workspace/api_rs.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"test/extended/testdata/apibinding/api_binding.yaml": testExtendedTestdataApibindingApi_bindingYaml,
	"test/extended/testdata/apibinding/api_export.yaml":  testExtendedTestdataApibindingApi_exportYaml,
	"test/extended/testdata/apibinding/api_rs.yaml":      testExtendedTestdataApibindingApi_rsYaml,
	"test/extended/testdata/apibinding/cowboy.yaml":      testExtendedTestdataApibindingCowboyYaml,
	"test/extended/testdata/apibinding/role_hack.yaml":   testExtendedTestdataApibindingRole_hackYaml,
	"test/extended/testdata/quota/quota.yaml":            testExtendedTestdataQuotaQuotaYaml,
	"test/extended/testdata/workspace/api_rs.yaml":       testExtendedTestdataWorkspaceApi_rsYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"test": {nil, map[string]*bintree{
		"extended": {nil, map[string]*bintree{
			"testdata": {nil, map[string]*bintree{
				"apibinding": {nil, map[string]*bintree{
					"api_binding.yaml": {testExtendedTestdataApibindingApi_bindingYaml, map[string]*bintree{}},
					"api_export.yaml":  {testExtendedTestdataApibindingApi_exportYaml, map[string]*bintree{}},
					"api_rs.yaml":      {testExtendedTestdataApibindingApi_rsYaml, map[string]*bintree{}},
					"cowboy.yaml":      {testExtendedTestdataApibindingCowboyYaml, map[string]*bintree{}},
					"role_hack.yaml":   {testExtendedTestdataApibindingRole_hackYaml, map[string]*bintree{}},
				}},
				"quota": {nil, map[string]*bintree{
					"quota.yaml": {testExtendedTestdataQuotaQuotaYaml, map[string]*bintree{}},
				}},
				"workspace": {nil, map[string]*bintree{
					"api_rs.yaml": {testExtendedTestdataWorkspaceApi_rsYaml, map[string]*bintree{}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
