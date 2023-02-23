package genoa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉拉CararaBarony struct {
	feud.BaseBarony
}

var BCarara卡拉拉 feud.Barony = &卡拉拉CararaBarony{}

func init() {
	f := BCarara卡拉拉.(*卡拉拉CararaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carara",
		TitleName: "卡拉拉",
		TitleCode: "b_carara",
	}
}
