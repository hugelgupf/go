// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binary

// +build 386 amd64 arm arm64 mipsle mips64le ppc64le

// NativeEndian is $GOARCH's implementation of byte order.
var NativeEndian = LittleEndian
