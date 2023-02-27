package monkh_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 臣赫尔TsenkherBarony struct {
	feud.BaseBarony
}

var BTsenkher臣赫尔 feud.Barony = &臣赫尔TsenkherBarony{}

func init() {
    f := BTsenkher臣赫尔.(*臣赫尔TsenkherBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsenkher",
		TitleName: "臣赫尔",
		TitleCode: "b_tsenkher",
	}
}
