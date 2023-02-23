package lucca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢卡LuccaBarony struct {
	feud.BaseBarony
}

var BLucca卢卡 feud.Barony = &卢卡LuccaBarony{}

func init() {
	f := BLucca卢卡.(*卢卡LuccaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lucca",
		TitleName: "卢卡",
		TitleCode: "b_lucca",
	}
}
