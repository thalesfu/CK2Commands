package ross

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丁沃尔DingwallBarony struct {
	feud.BaseBarony
}

var BDingwall丁沃尔 feud.Barony = &丁沃尔DingwallBarony{}

func init() {
	f := BDingwall丁沃尔.(*丁沃尔DingwallBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dingwall",
		TitleName: "丁沃尔",
		TitleCode: "b_dingwall",
	}
}
