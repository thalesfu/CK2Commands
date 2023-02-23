package reval

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡莱万RevalBarony struct {
	feud.BaseBarony
}

var BReval卡莱万 feud.Barony = &卡莱万RevalBarony{}

func init() {
	f := BReval卡莱万.(*卡莱万RevalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reval",
		TitleName: "卡莱万",
		TitleCode: "b_reval",
	}
}
