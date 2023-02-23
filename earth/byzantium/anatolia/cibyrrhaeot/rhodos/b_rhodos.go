package rhodos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗德岛RhodosBarony struct {
	feud.BaseBarony
}

var BRhodos罗德岛 feud.Barony = &罗德岛RhodosBarony{}

func init() {
	f := BRhodos罗德岛.(*罗德岛RhodosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rhodos",
		TitleName: "罗德岛",
		TitleCode: "b_rhodos",
	}
}
