package drutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达什科夫卡DashkovkaBarony struct {
	feud.BaseBarony
}

var BDashkovka达什科夫卡 feud.Barony = &达什科夫卡DashkovkaBarony{}

func init() {
	f := BDashkovka达什科夫卡.(*达什科夫卡DashkovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dashkovka",
		TitleName: "达什科夫卡",
		TitleCode: "b_dashkovka",
	}
}
