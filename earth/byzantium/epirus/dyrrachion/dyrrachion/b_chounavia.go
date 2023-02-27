package dyrrachion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库纳维亚ChounaviaBarony struct {
	feud.BaseBarony
}

var BChounavia库纳维亚 feud.Barony = &库纳维亚ChounaviaBarony{}

func init() {
    f := BChounavia库纳维亚.(*库纳维亚ChounaviaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chounavia",
		TitleName: "库纳维亚",
		TitleCode: "b_chounavia",
	}
}
