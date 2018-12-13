// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binary

// +build mips mips64 ppc64 s390x

// NativeEndian is $GOARCH's implementation of byte order.
var NativeEndian = BigEndian
