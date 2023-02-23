package shetland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 廷沃尔TingwallBarony struct {
	feud.BaseBarony
}

var BTingwall廷沃尔 feud.Barony = &廷沃尔TingwallBarony{}

func init() {
	f := BTingwall廷沃尔.(*廷沃尔TingwallBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tingwall",
		TitleName: "廷沃尔",
		TitleCode: "b_tingwall",
	}
}
