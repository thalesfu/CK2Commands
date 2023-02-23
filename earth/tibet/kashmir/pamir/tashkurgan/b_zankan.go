package tashkurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赞坎ZankanBarony struct {
	feud.BaseBarony
}

var BZankan赞坎 feud.Barony = &赞坎ZankanBarony{}

func init() {
	f := BZankan赞坎.(*赞坎ZankanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zankan",
		TitleName: "赞坎",
		TitleCode: "b_zankan",
	}
}
