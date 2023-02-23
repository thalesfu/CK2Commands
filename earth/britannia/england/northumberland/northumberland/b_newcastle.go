package northumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纽卡斯尔NewcastleBarony struct {
	feud.BaseBarony
}

var BNewcastle纽卡斯尔 feud.Barony = &纽卡斯尔NewcastleBarony{}

func init() {
	f := BNewcastle纽卡斯尔.(*纽卡斯尔NewcastleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "newcastle",
		TitleName: "纽卡斯尔",
		TitleCode: "b_newcastle",
	}
}
