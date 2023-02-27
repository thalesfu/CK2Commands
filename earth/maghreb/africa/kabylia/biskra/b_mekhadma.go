package biskra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅克哈德马MekhadmaBarony struct {
	feud.BaseBarony
}

var BMekhadma梅克哈德马 feud.Barony = &梅克哈德马MekhadmaBarony{}

func init() {
    f := BMekhadma梅克哈德马.(*梅克哈德马MekhadmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mekhadma",
		TitleName: "梅克哈德马",
		TitleCode: "b_mekhadma",
	}
}
