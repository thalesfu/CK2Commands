package snassen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾赫菲尔AhfirBarony struct {
	feud.BaseBarony
}

var BAhfir艾赫菲尔 feud.Barony = &艾赫菲尔AhfirBarony{}

func init() {
	f := BAhfir艾赫菲尔.(*艾赫菲尔AhfirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahfir",
		TitleName: "艾赫菲尔",
		TitleCode: "b_ahfir",
	}
}
