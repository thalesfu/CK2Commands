package temes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布齐亚什BuziasBarony struct {
	feud.BaseBarony
}

var BBuzias布齐亚什 feud.Barony = &布齐亚什BuziasBarony{}

func init() {
	f := BBuzias布齐亚什.(*布齐亚什BuziasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buzias",
		TitleName: "布齐亚什",
		TitleCode: "b_buzias",
	}
}
