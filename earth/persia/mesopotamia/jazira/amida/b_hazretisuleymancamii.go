package amida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈兹拉特苏莱曼寺HazretisuleymancamiiBarony struct {
	feud.BaseBarony
}

var BHazretisuleymancamii哈兹拉特苏莱曼寺 feud.Barony = &哈兹拉特苏莱曼寺HazretisuleymancamiiBarony{}

func init() {
    f := BHazretisuleymancamii哈兹拉特苏莱曼寺.(*哈兹拉特苏莱曼寺HazretisuleymancamiiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hazretisuleymancamii",
		TitleName: "哈兹拉特苏莱曼寺",
		TitleCode: "b_hazretisuleymancamii",
	}
}
