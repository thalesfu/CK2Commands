package srirangapatna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩索罗MysoreBarony struct {
	feud.BaseBarony
}

var BMysore摩索罗 feud.Barony = &摩索罗MysoreBarony{}

func init() {
	f := BMysore摩索罗.(*摩索罗MysoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mysore",
		TitleName: "摩索罗",
		TitleCode: "b_mysore",
	}
}
