package bhakkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尸佉罗补罗ShikarpurBarony struct {
	feud.BaseBarony
}

var BShikarpur尸佉罗补罗 feud.Barony = &尸佉罗补罗ShikarpurBarony{}

func init() {
	f := BShikarpur尸佉罗补罗.(*尸佉罗补罗ShikarpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shikarpur",
		TitleName: "尸佉罗补罗",
		TitleCode: "b_shikarpur",
	}
}
