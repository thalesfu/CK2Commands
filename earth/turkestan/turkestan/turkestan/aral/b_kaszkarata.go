package aral

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡什卡拉塔KaszkarataBarony struct {
	feud.BaseBarony
}

var BKaszkarata卡什卡拉塔 feud.Barony = &卡什卡拉塔KaszkarataBarony{}

func init() {
    f := BKaszkarata卡什卡拉塔.(*卡什卡拉塔KaszkarataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaszkarata",
		TitleName: "卡什卡拉塔",
		TitleCode: "b_kaszkarata",
	}
}
