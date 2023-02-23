package viken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫德HedeBarony struct {
	feud.BaseBarony
}

var BHede赫德 feud.Barony = &赫德HedeBarony{}

func init() {
	f := BHede赫德.(*赫德HedeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hede",
		TitleName: "赫德",
		TitleCode: "b_hede",
	}
}
