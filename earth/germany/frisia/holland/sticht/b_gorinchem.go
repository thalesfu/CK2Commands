package sticht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍林赫姆GorinchemBarony struct {
	feud.BaseBarony
}

var BGorinchem霍林赫姆 feud.Barony = &霍林赫姆GorinchemBarony{}

func init() {
	f := BGorinchem霍林赫姆.(*霍林赫姆GorinchemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorinchem",
		TitleName: "霍林赫姆",
		TitleCode: "b_gorinchem",
	}
}
