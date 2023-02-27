package starodub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔罗杜布StarodubBarony struct {
	feud.BaseBarony
}

var BStarodub斯塔罗杜布 feud.Barony = &斯塔罗杜布StarodubBarony{}

func init() {
    f := BStarodub斯塔罗杜布.(*斯塔罗杜布StarodubBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "starodub",
		TitleName: "斯塔罗杜布",
		TitleCode: "b_starodub",
	}
}
