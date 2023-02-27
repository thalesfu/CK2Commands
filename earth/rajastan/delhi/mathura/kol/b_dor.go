package kol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兜尔DorBarony struct {
	feud.BaseBarony
}

var BDor兜尔 feud.Barony = &兜尔DorBarony{}

func init() {
    f := BDor兜尔.(*兜尔DorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dor",
		TitleName: "兜尔",
		TitleCode: "b_dor",
	}
}
