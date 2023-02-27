package zavarot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安杰格AndegBarony struct {
	feud.BaseBarony
}

var BAndeg安杰格 feud.Barony = &安杰格AndegBarony{}

func init() {
    f := BAndeg安杰格.(*安杰格AndegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andeg",
		TitleName: "安杰格",
		TitleCode: "b_andeg",
	}
}
