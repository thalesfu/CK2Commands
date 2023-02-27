package charkliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毛家乡MaojiaxiangBarony struct {
	feud.BaseBarony
}

var BMaojiaxiang毛家乡 feud.Barony = &毛家乡MaojiaxiangBarony{}

func init() {
    f := BMaojiaxiang毛家乡.(*毛家乡MaojiaxiangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maojiaxiang",
		TitleName: "毛家乡",
		TitleCode: "b_maojiaxiang",
	}
}
