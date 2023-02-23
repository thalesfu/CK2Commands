package monreal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴兰ParanBarony struct {
	feud.BaseBarony
}

var BParan巴兰 feud.Barony = &巴兰ParanBarony{}

func init() {
	f := BParan巴兰.(*巴兰ParanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paran",
		TitleName: "巴兰",
		TitleCode: "b_paran",
	}
}
