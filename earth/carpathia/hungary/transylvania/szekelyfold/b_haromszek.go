package szekelyfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈罗姆塞克HaromszekBarony struct {
	feud.BaseBarony
}

var BHaromszek哈罗姆塞克 feud.Barony = &哈罗姆塞克HaromszekBarony{}

func init() {
	f := BHaromszek哈罗姆塞克.(*哈罗姆塞克HaromszekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haromszek",
		TitleName: "哈罗姆塞克",
		TitleCode: "b_haromszek",
	}
}
