package gaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍夷罗迦提SherghatiBarony struct {
	feud.BaseBarony
}

var BSherghati舍夷罗迦提 feud.Barony = &舍夷罗迦提SherghatiBarony{}

func init() {
	f := BSherghati舍夷罗迦提.(*舍夷罗迦提SherghatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sherghati",
		TitleName: "舍夷罗迦提",
		TitleCode: "b_sherghati",
	}
}
