package meknes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂尔希拉TerhirhaBarony struct {
	feud.BaseBarony
}

var BTerhirha蒂尔希拉 feud.Barony = &蒂尔希拉TerhirhaBarony{}

func init() {
	f := BTerhirha蒂尔希拉.(*蒂尔希拉TerhirhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terhirha",
		TitleName: "蒂尔希拉",
		TitleCode: "b_terhirha",
	}
}
