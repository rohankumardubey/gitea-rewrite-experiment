// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(http.Dir("../../public"), vfsgen.Options{
		PackageName:  "assets",
		BuildTags:    "bindata",
		VariableName: "assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
