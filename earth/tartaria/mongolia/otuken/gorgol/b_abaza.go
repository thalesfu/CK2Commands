package gorgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴扎AbazaBarony struct {
	feud.BaseBarony
}

var BAbaza阿巴扎 feud.Barony = &阿巴扎AbazaBarony{}

func init() {
    f := BAbaza阿巴扎.(*阿巴扎AbazaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abaza",
		TitleName: "阿巴扎",
		TitleCode: "b_abaza",
	}
}
