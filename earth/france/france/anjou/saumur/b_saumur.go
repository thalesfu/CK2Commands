package saumur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索米尔SaumurBarony struct {
	feud.BaseBarony
}

var BSaumur索米尔 feud.Barony = &索米尔SaumurBarony{}

func init() {
	f := BSaumur索米尔.(*索米尔SaumurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saumur",
		TitleName: "索米尔",
		TitleCode: "b_saumur",
	}
}
