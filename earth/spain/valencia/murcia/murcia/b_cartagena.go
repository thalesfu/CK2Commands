package murcia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塔赫纳CartagenaBarony struct {
	feud.BaseBarony
}

var BCartagena卡塔赫纳 feud.Barony = &卡塔赫纳CartagenaBarony{}

func init() {
	f := BCartagena卡塔赫纳.(*卡塔赫纳CartagenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cartagena",
		TitleName: "卡塔赫纳",
		TitleCode: "b_cartagena",
	}
}
