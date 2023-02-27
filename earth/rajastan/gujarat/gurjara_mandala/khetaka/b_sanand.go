package khetaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑恩德SanandBarony struct {
	feud.BaseBarony
}

var BSanand桑恩德 feud.Barony = &桑恩德SanandBarony{}

func init() {
    f := BSanand桑恩德.(*桑恩德SanandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanand",
		TitleName: "桑恩德",
		TitleCode: "b_sanand",
	}
}
