package tver

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特维尔TverBarony struct {
	feud.BaseBarony
}

var BTver特维尔 feud.Barony = &特维尔TverBarony{}

func init() {
    f := BTver特维尔.(*特维尔TverBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tver",
		TitleName: "特维尔",
		TitleCode: "b_tver",
	}
}
