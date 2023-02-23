package yarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麻札MazarBarony struct {
	feud.BaseBarony
}

var BMazar麻札 feud.Barony = &麻札MazarBarony{}

func init() {
	f := BMazar麻札.(*麻札MazarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazar",
		TitleName: "麻札",
		TitleCode: "b_mazar",
	}
}
