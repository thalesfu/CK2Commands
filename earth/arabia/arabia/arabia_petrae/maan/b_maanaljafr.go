package maan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰夫尔MaanaljafrBarony struct {
	feud.BaseBarony
}

var BMaanaljafr杰夫尔 feud.Barony = &杰夫尔MaanaljafrBarony{}

func init() {
    f := BMaanaljafr杰夫尔.(*杰夫尔MaanaljafrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maanaljafr",
		TitleName: "杰夫尔",
		TitleCode: "b_maanaljafr",
	}
}
