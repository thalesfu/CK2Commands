package maroneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑德雅XantheiaBarony struct {
	feud.BaseBarony
}

var BXantheia桑德雅 feud.Barony = &桑德雅XantheiaBarony{}

func init() {
	f := BXantheia桑德雅.(*桑德雅XantheiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xantheia",
		TitleName: "桑德雅",
		TitleCode: "b_xantheia",
	}
}
