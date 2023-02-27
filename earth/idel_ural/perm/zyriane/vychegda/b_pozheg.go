package vychegda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波热格PozhegBarony struct {
	feud.BaseBarony
}

var BPozheg波热格 feud.Barony = &波热格PozhegBarony{}

func init() {
    f := BPozheg波热格.(*波热格PozhegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pozheg",
		TitleName: "波热格",
		TitleCode: "b_pozheg",
	}
}
