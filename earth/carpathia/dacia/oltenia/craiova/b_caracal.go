package craiova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉卡尔CaracalBarony struct {
	feud.BaseBarony
}

var BCaracal卡拉卡尔 feud.Barony = &卡拉卡尔CaracalBarony{}

func init() {
	f := BCaracal卡拉卡尔.(*卡拉卡尔CaracalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caracal",
		TitleName: "卡拉卡尔",
		TitleCode: "b_caracal",
	}
}
