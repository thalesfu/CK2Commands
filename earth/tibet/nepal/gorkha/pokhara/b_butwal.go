package pokhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布德沃尔ButwalBarony struct {
	feud.BaseBarony
}

var BButwal布德沃尔 feud.Barony = &布德沃尔ButwalBarony{}

func init() {
	f := BButwal布德沃尔.(*布德沃尔ButwalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "butwal",
		TitleName: "布德沃尔",
		TitleCode: "b_butwal",
	}
}
