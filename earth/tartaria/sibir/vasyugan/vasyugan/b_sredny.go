package vasyugan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯列德尼SrednyBarony struct {
	feud.BaseBarony
}

var BSredny斯列德尼 feud.Barony = &斯列德尼SrednyBarony{}

func init() {
    f := BSredny斯列德尼.(*斯列德尼SrednyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sredny",
		TitleName: "斯列德尼",
		TitleCode: "b_sredny",
	}
}
