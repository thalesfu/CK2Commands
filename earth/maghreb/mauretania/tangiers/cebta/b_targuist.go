package cebta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔吉斯特TarguistBarony struct {
	feud.BaseBarony
}

var BTarguist塔尔吉斯特 feud.Barony = &塔尔吉斯特TarguistBarony{}

func init() {
	f := BTarguist塔尔吉斯特.(*塔尔吉斯特TarguistBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "targuist",
		TitleName: "塔尔吉斯特",
		TitleCode: "b_targuist",
	}
}
