package tura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图拉TuraBarony struct {
	feud.BaseBarony
}

var BTura图拉 feud.Barony = &图拉TuraBarony{}

func init() {
	f := BTura图拉.(*图拉TuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tura",
		TitleName: "图拉",
		TitleCode: "b_tura",
	}
}
