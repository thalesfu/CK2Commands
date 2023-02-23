package lhoyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安尼尼AniniBarony struct {
	feud.BaseBarony
}

var BAnini安尼尼 feud.Barony = &安尼尼AniniBarony{}

func init() {
	f := BAnini安尼尼.(*安尼尼AniniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anini",
		TitleName: "安尼尼",
		TitleCode: "b_anini",
	}
}
