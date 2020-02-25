// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"testing"
)

func Test_ReadWriter(t *testing.T) {
	rw, err := NewReadWriter(`test.zip`)
	if err != nil {
		t.Error(err)
	}
	defer rw.Close()
	if err := rw.Add(`test123321`, []byte(`hello world test`)); err != nil {
		t.Fatal(err)
	}
	r, err := OpenReader(`test.zip`)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range r.File {
		if f.Name == `test123321` {
			t.Log(`Adding file: SUCCESS`)
			t.SkipNow()
		}
	}
	t.Fatal(`no file added`)
}
