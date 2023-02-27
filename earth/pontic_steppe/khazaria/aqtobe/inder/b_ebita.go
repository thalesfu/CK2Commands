package inder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃比德EbitaBarony struct {
	feud.BaseBarony
}

var BEbita埃比德 feud.Barony = &埃比德EbitaBarony{}

func init() {
    f := BEbita埃比德.(*埃比德EbitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ebita",
		TitleName: "埃比德",
		TitleCode: "b_ebita",
	}
}
