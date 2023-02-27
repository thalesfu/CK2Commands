package aqtobe

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔_库特尔塔斯AulkutyrtasBarony struct {
	feud.BaseBarony
}

var BAulkutyrtas奥尔_库特尔塔斯 feud.Barony = &奥尔_库特尔塔斯AulkutyrtasBarony{}

func init() {
    f := BAulkutyrtas奥尔_库特尔塔斯.(*奥尔_库特尔塔斯AulkutyrtasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aulkutyrtas",
		TitleName: "奥尔_库特尔塔斯",
		TitleCode: "b_aulkutyrtas",
	}
}
