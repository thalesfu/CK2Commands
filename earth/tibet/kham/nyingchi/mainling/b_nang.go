package mainling

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗县NangBarony struct {
	feud.BaseBarony
}

var BNang朗县 feud.Barony = &朗县NangBarony{}

func init() {
	f := BNang朗县.(*朗县NangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nang",
		TitleName: "朗县",
		TitleCode: "b_nang",
	}
}
