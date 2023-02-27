package kurs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅泽MedzeBarony struct {
	feud.BaseBarony
}

var BMedze梅泽 feud.Barony = &梅泽MedzeBarony{}

func init() {
    f := BMedze梅泽.(*梅泽MedzeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medze",
		TitleName: "梅泽",
		TitleCode: "b_medze",
	}
}
