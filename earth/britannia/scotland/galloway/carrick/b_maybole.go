package carrick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅博尔MayboleBarony struct {
	feud.BaseBarony
}

var BMaybole梅博尔 feud.Barony = &梅博尔MayboleBarony{}

func init() {
	f := BMaybole梅博尔.(*梅博尔MayboleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maybole",
		TitleName: "梅博尔",
		TitleCode: "b_maybole",
	}
}
