package yatenga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪雷拉贝DirelabeBarony struct {
	feud.BaseBarony
}

var BDirelabe迪雷拉贝 feud.Barony = &迪雷拉贝DirelabeBarony{}

func init() {
    f := BDirelabe迪雷拉贝.(*迪雷拉贝DirelabeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "direlabe",
		TitleName: "迪雷拉贝",
		TitleCode: "b_direlabe",
	}
}
