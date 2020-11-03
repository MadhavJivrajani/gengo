// +build !ignore_autogenerated

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

// Code generated by defaulter-gen. DO NOT EDIT.

package marker

import (
	"encoding/json"
	"reflect"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&Defaulted{}, func(obj interface{}) { SetObjectDefaults_Defaulted(obj.(*Defaulted)) })
	scheme.AddTypeDefaultingFunc(&SubStruct{}, func(obj interface{}) { SetObjectDefaults_SubStruct(obj.(*SubStruct)) })
	return nil
}

func SetObjectDefaults_Defaulted(in *Defaulted) {
	if reflect.ValueOf(in.Field).IsZero() {
		if err := json.Unmarshal([]byte(`"bar"`), &in.Field); err != nil {
			panic(err)
		}
	}
	if reflect.ValueOf(in.OtherField).IsZero() {
		if err := json.Unmarshal([]byte(`0`), &in.OtherField); err != nil {
			panic(err)
		}
	}
	if reflect.ValueOf(in.List).IsZero() {
		if err := json.Unmarshal([]byte(`["foo", "bar"]`), &in.List); err != nil {
			panic(err)
		}
	}
	for i := range in.List {
		if reflect.ValueOf(in.List[i]).IsZero() {
			if err := json.Unmarshal([]byte(`"apple"`), &in.List[i]); err != nil {
				panic(err)
			}
		}
	}
	if reflect.ValueOf(in.Sub).IsZero() {
		if err := json.Unmarshal([]byte(`{"s": "foo", "i": 5}`), &in.Sub); err != nil {
			panic(err)
		}
	}
	if in.Sub != nil {
		SetObjectDefaults_SubStruct(in.Sub)
	}
	SetObjectDefaults_SubStruct(&in.OtherSub)
	if reflect.ValueOf(in.Map).IsZero() {
		if err := json.Unmarshal([]byte(`{"foo": "bar"}`), &in.Map); err != nil {
			panic(err)
		}
	}
	for i_Map := range in.Map {
		if reflect.ValueOf(in.Map[i_Map]).IsZero() {
			i_Map_default := in.Map[i_Map]
			if err := json.Unmarshal([]byte(`"apple"`), &i_Map_default); err != nil {
				panic(err)
			}
			in.Map[i_Map] = i_Map_default
		}
	}
}

func SetObjectDefaults_SubStruct(in *SubStruct) {
	if reflect.ValueOf(in.I).IsZero() {
		if err := json.Unmarshal([]byte(`1`), &in.I); err != nil {
			panic(err)
		}
	}
}
