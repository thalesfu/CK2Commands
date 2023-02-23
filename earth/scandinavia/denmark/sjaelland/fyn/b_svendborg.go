package fyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯文堡SvendborgBarony struct {
	feud.BaseBarony
}

var BSvendborg斯文堡 feud.Barony = &斯文堡SvendborgBarony{}

func init() {
	f := BSvendborg斯文堡.(*斯文堡SvendborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svendborg",
		TitleName: "斯文堡",
		TitleCode: "b_svendborg",
	}
}
