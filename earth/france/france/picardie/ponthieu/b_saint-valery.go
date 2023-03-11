package ponthieu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣瓦莱里Saint_valeryBarony struct {
	feud.BaseBarony
}

var BSaint_valery圣瓦莱里 feud.Barony = &圣瓦莱里Saint_valeryBarony{}

func init() {
    f := BSaint_valery圣瓦莱里.(*圣瓦莱里Saint_valeryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saint_valery",
		TitleName: "圣瓦莱里",
		TitleCode: "b_saint-valery",
	}
}
