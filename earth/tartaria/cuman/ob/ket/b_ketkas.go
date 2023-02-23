package ket

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 憩加斯KetkasBarony struct {
	feud.BaseBarony
}

var BKetkas憩加斯 feud.Barony = &憩加斯KetkasBarony{}

func init() {
	f := BKetkas憩加斯.(*憩加斯KetkasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ketkas",
		TitleName: "憩加斯",
		TitleCode: "b_ketkas",
	}
}
