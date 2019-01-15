// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main // import "go.zchee.io/cmd/manpagejp"

import (
	"log"
	"os"

	manpagejp "go.zchee.io/cmd/pkg/manpagejp"
)

func main() {
	jmp := manpagejp.NewJMProject()

	if err := jmp.Download(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	if err := jmp.Extract(); err != nil {
		log.Fatal(err)
	}

	if err := jmp.Copy(os.Args[2], os.Args[1]); err != nil {
		log.Fatal(err)
	}
}