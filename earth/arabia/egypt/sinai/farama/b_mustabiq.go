package farama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆斯塔比克MustabiqBarony struct {
	feud.BaseBarony
}

var BMustabiq穆斯塔比克 feud.Barony = &穆斯塔比克MustabiqBarony{}

func init() {
	f := BMustabiq穆斯塔比克.(*穆斯塔比克MustabiqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mustabiq",
		TitleName: "穆斯塔比克",
		TitleCode: "b_mustabiq",
	}
}
