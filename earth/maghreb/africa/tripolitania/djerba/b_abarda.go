package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴拉达AbardaBarony struct {
	feud.BaseBarony
}

var BAbarda阿巴拉达 feud.Barony = &阿巴拉达AbardaBarony{}

func init() {
    f := BAbarda阿巴拉达.(*阿巴拉达AbardaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abarda",
		TitleName: "阿巴拉达",
		TitleCode: "b_abarda",
	}
}
