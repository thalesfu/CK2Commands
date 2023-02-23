package fyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼堡NyborgBarony struct {
	feud.BaseBarony
}

var BNyborg尼堡 feud.Barony = &尼堡NyborgBarony{}

func init() {
	f := BNyborg尼堡.(*尼堡NyborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyborg",
		TitleName: "尼堡",
		TitleCode: "b_nyborg",
	}
}
