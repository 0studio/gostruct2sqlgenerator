// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagParsing(t *testing.T) {
	fields := parseTag("field,foobar,foo")
	assert.Equal(t, "field", fields[0].TagKey)
	assert.Equal(t, "foobar", fields[1].TagKey)
	assert.Equal(t, "foo", fields[2].TagKey)
}
