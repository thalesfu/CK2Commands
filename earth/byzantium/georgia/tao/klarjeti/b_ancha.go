package klarjeti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安查AnchaBarony struct {
	feud.BaseBarony
}

var BAncha安查 feud.Barony = &安查AnchaBarony{}

func init() {
    f := BAncha安查.(*安查AnchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ancha",
		TitleName: "安查",
		TitleCode: "b_ancha",
	}
}
