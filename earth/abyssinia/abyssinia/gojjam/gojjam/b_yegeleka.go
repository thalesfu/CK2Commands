package gojjam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶格勒卡YegelekaBarony struct {
	feud.BaseBarony
}

var BYegeleka耶格勒卡 feud.Barony = &耶格勒卡YegelekaBarony{}

func init() {
	f := BYegeleka耶格勒卡.(*耶格勒卡YegelekaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yegeleka",
		TitleName: "耶格勒卡",
		TitleCode: "b_yegeleka",
	}
}
