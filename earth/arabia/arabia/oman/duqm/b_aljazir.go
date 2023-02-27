package duqm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰济拉AljazirBarony struct {
	feud.BaseBarony
}

var BAljazir杰济拉 feud.Barony = &杰济拉AljazirBarony{}

func init() {
    f := BAljazir杰济拉.(*杰济拉AljazirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aljazir",
		TitleName: "杰济拉",
		TitleCode: "b_aljazir",
	}
}
