package kimak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷努尔TenulBarony struct {
	feud.BaseBarony
}

var BTenul捷努尔 feud.Barony = &捷努尔TenulBarony{}

func init() {
	f := BTenul捷努尔.(*捷努尔TenulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tenul",
		TitleName: "捷努尔",
		TitleCode: "b_tenul",
	}
}
