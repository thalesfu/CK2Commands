package laukaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯乌鲁KeuruuBarony struct {
	feud.BaseBarony
}

var BKeuruu凯乌鲁 feud.Barony = &凯乌鲁KeuruuBarony{}

func init() {
	f := BKeuruu凯乌鲁.(*凯乌鲁KeuruuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keuruu",
		TitleName: "凯乌鲁",
		TitleCode: "b_keuruu",
	}
}
