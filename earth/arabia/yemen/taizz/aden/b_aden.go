package aden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚丁AdenBarony struct {
	feud.BaseBarony
}

var BAden亚丁 feud.Barony = &亚丁AdenBarony{}

func init() {
	f := BAden亚丁.(*亚丁AdenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aden",
		TitleName: "亚丁",
		TitleCode: "b_aden",
	}
}
