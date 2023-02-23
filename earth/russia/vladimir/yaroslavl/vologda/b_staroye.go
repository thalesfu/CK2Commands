package vologda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔罗耶StaroyeBarony struct {
	feud.BaseBarony
}

var BStaroye斯塔罗耶 feud.Barony = &斯塔罗耶StaroyeBarony{}

func init() {
	f := BStaroye斯塔罗耶.(*斯塔罗耶StaroyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "staroye",
		TitleName: "斯塔罗耶",
		TitleCode: "b_staroye",
	}
}
