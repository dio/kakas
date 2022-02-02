// Copyright 2022 Dhi Aurrahman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"testing"

	"github.com/dio/kakas/test/generated"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestGenerated(t *testing.T) {
	a := &generated.TagType{FieldA: 1, FieldB: "a"}
	b := &generated.TagType{FieldA: 2, FieldB: "b"}
	a.DeepCopyInto(b)
	require.True(t, cmp.Equal(a, b, protocmp.Transform()))
	bytes, err := a.MarshalJSON()
	require.NoError(t, err)
	require.NoError(t, a.UnmarshalJSON(bytes))
	require.True(t, cmp.Equal(a, b, protocmp.Transform()))

	c := &generated.PackageSent{}
	d := &generated.PackageReceived{}
	require.Equal(t, c.ProtoReflect().Descriptor().FullName(), protoreflect.FullName(d.EventMetadata().PreviousTypeUrls[0]))
}
