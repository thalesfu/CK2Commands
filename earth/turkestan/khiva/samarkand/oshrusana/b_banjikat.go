package oshrusana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班吉卡特BanjikatBarony struct {
	feud.BaseBarony
}

var BBanjikat班吉卡特 feud.Barony = &班吉卡特BanjikatBarony{}

func init() {
    f := BBanjikat班吉卡特.(*班吉卡特BanjikatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banjikat",
		TitleName: "班吉卡特",
		TitleCode: "b_banjikat",
	}
}
