package khantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别廖佐沃BerezovoBarony struct {
	feud.BaseBarony
}

var BBerezovo别廖佐沃 feud.Barony = &别廖佐沃BerezovoBarony{}

func init() {
    f := BBerezovo别廖佐沃.(*别廖佐沃BerezovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berezovo",
		TitleName: "别廖佐沃",
		TitleCode: "b_berezovo",
	}
}
