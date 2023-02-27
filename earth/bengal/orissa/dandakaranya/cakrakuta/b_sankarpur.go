package cakrakuta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 商羯罗补罗SankarpurBarony struct {
	feud.BaseBarony
}

var BSankarpur商羯罗补罗 feud.Barony = &商羯罗补罗SankarpurBarony{}

func init() {
    f := BSankarpur商羯罗补罗.(*商羯罗补罗SankarpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sankarpur",
		TitleName: "商羯罗补罗",
		TitleCode: "b_sankarpur",
	}
}
