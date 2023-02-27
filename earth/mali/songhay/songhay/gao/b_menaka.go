package gao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅纳卡MenakaBarony struct {
	feud.BaseBarony
}

var BMenaka梅纳卡 feud.Barony = &梅纳卡MenakaBarony{}

func init() {
    f := BMenaka梅纳卡.(*梅纳卡MenakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "menaka",
		TitleName: "梅纳卡",
		TitleCode: "b_menaka",
	}
}
