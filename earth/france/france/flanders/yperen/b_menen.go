package yperen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅嫩MenenBarony struct {
	feud.BaseBarony
}

var BMenen梅嫩 feud.Barony = &梅嫩MenenBarony{}

func init() {
    f := BMenen梅嫩.(*梅嫩MenenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "menen",
		TitleName: "梅嫩",
		TitleCode: "b_menen",
	}
}
