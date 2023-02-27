package kolomna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图拉TulaBarony struct {
	feud.BaseBarony
}

var BTula图拉 feud.Barony = &图拉TulaBarony{}

func init() {
    f := BTula图拉.(*图拉TulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tula",
		TitleName: "图拉",
		TitleCode: "b_tula",
	}
}
