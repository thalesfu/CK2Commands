package koshma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡赖HuraiBarony struct {
	feud.BaseBarony
}

var BHurai胡赖 feud.Barony = &胡赖HuraiBarony{}

func init() {
    f := BHurai胡赖.(*胡赖HuraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hurai",
		TitleName: "胡赖",
		TitleCode: "b_hurai",
	}
}
