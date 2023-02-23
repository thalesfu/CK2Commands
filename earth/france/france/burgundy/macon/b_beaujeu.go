package macon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博热BeaujeuBarony struct {
	feud.BaseBarony
}

var BBeaujeu博热 feud.Barony = &博热BeaujeuBarony{}

func init() {
	f := BBeaujeu博热.(*博热BeaujeuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beaujeu",
		TitleName: "博热",
		TitleCode: "b_beaujeu",
	}
}
