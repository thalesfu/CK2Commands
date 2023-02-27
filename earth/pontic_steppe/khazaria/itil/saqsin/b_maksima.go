package saqsin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马克西马MaksimaBarony struct {
	feud.BaseBarony
}

var BMaksima马克西马 feud.Barony = &马克西马MaksimaBarony{}

func init() {
    f := BMaksima马克西马.(*马克西马MaksimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maksima",
		TitleName: "马克西马",
		TitleCode: "b_maksima",
	}
}
