package dulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夏日哈XiarihaBarony struct {
	feud.BaseBarony
}

var BXiariha夏日哈 feud.Barony = &夏日哈XiarihaBarony{}

func init() {
	f := BXiariha夏日哈.(*夏日哈XiarihaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiariha",
		TitleName: "夏日哈",
		TitleCode: "b_xiariha",
	}
}
