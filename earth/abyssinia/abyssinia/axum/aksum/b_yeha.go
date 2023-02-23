package aksum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶哈YehaBarony struct {
	feud.BaseBarony
}

var BYeha耶哈 feud.Barony = &耶哈YehaBarony{}

func init() {
	f := BYeha耶哈.(*耶哈YehaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yeha",
		TitleName: "耶哈",
		TitleCode: "b_yeha",
	}
}
