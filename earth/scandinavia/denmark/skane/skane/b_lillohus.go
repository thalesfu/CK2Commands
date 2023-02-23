package skane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利勒胡斯LillohusBarony struct {
	feud.BaseBarony
}

var BLillohus利勒胡斯 feud.Barony = &利勒胡斯LillohusBarony{}

func init() {
	f := BLillohus利勒胡斯.(*利勒胡斯LillohusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lillohus",
		TitleName: "利勒胡斯",
		TitleCode: "b_lillohus",
	}
}
