package tyrone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓甘嫩DungannonBarony struct {
	feud.BaseBarony
}

var BDungannon邓甘嫩 feud.Barony = &邓甘嫩DungannonBarony{}

func init() {
    f := BDungannon邓甘嫩.(*邓甘嫩DungannonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dungannon",
		TitleName: "邓甘嫩",
		TitleCode: "b_dungannon",
	}
}
