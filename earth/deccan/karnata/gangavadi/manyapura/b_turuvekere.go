package manyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图鲁韦凯雷TuruvekereBarony struct {
	feud.BaseBarony
}

var BTuruvekere图鲁韦凯雷 feud.Barony = &图鲁韦凯雷TuruvekereBarony{}

func init() {
    f := BTuruvekere图鲁韦凯雷.(*图鲁韦凯雷TuruvekereBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turuvekere",
		TitleName: "图鲁韦凯雷",
		TitleCode: "b_turuvekere",
	}
}
