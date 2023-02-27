package maan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈赛MaanalhasaBarony struct {
	feud.BaseBarony
}

var BMaanalhasa哈赛 feud.Barony = &哈赛MaanalhasaBarony{}

func init() {
    f := BMaanalhasa哈赛.(*哈赛MaanalhasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maanalhasa",
		TitleName: "哈赛",
		TitleCode: "b_maanalhasa",
	}
}
