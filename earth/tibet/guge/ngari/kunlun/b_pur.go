package kunlun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普尔PurBarony struct {
	feud.BaseBarony
}

var BPur普尔 feud.Barony = &普尔PurBarony{}

func init() {
	f := BPur普尔.(*普尔PurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pur",
		TitleName: "普尔",
		TitleCode: "b_pur",
	}
}
