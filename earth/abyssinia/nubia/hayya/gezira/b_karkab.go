package gezira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔卡布KarkabBarony struct {
	feud.BaseBarony
}

var BKarkab卡尔卡布 feud.Barony = &卡尔卡布KarkabBarony{}

func init() {
	f := BKarkab卡尔卡布.(*卡尔卡布KarkabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karkab",
		TitleName: "卡尔卡布",
		TitleCode: "b_karkab",
	}
}
