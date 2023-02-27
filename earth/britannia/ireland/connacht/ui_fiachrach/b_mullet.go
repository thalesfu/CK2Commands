package ui_fiachrach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆雷特MulletBarony struct {
	feud.BaseBarony
}

var BMullet穆雷特 feud.Barony = &穆雷特MulletBarony{}

func init() {
    f := BMullet穆雷特.(*穆雷特MulletBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mullet",
		TitleName: "穆雷特",
		TitleCode: "b_mullet",
	}
}
