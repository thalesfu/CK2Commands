package kildare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康斯喀林KnockaulinBarony struct {
	feud.BaseBarony
}

var BKnockaulin康斯喀林 feud.Barony = &康斯喀林KnockaulinBarony{}

func init() {
	f := BKnockaulin康斯喀林.(*康斯喀林KnockaulinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knockaulin",
		TitleName: "康斯喀林",
		TitleCode: "b_knockaulin",
	}
}
