package kharkhiraa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 车车格TsetsegBarony struct {
	feud.BaseBarony
}

var BTsetseg车车格 feud.Barony = &车车格TsetsegBarony{}

func init() {
    f := BTsetseg车车格.(*车车格TsetsegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsetseg",
		TitleName: "车车格",
		TitleCode: "b_tsetseg",
	}
}
