package sinjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈特拉HatraBarony struct {
	feud.BaseBarony
}

var BHatra哈特拉 feud.Barony = &哈特拉HatraBarony{}

func init() {
	f := BHatra哈特拉.(*哈特拉HatraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hatra",
		TitleName: "哈特拉",
		TitleCode: "b_hatra",
	}
}
