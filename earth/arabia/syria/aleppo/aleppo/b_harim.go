package aleppo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈里姆HarimBarony struct {
	feud.BaseBarony
}

var BHarim哈里姆 feud.Barony = &哈里姆HarimBarony{}

func init() {
	f := BHarim哈里姆.(*哈里姆HarimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harim",
		TitleName: "哈里姆",
		TitleCode: "b_harim",
	}
}
