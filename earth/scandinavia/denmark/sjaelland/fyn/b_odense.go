package fyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧登塞OdenseBarony struct {
	feud.BaseBarony
}

var BOdense欧登塞 feud.Barony = &欧登塞OdenseBarony{}

func init() {
	f := BOdense欧登塞.(*欧登塞OdenseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odense",
		TitleName: "欧登塞",
		TitleCode: "b_odense",
	}
}
