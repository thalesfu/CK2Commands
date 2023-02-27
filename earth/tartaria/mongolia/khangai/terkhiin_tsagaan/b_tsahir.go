package terkhiin_tsagaan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查黑尔TsahirBarony struct {
	feud.BaseBarony
}

var BTsahir查黑尔 feud.Barony = &查黑尔TsahirBarony{}

func init() {
    f := BTsahir查黑尔.(*查黑尔TsahirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsahir",
		TitleName: "查黑尔",
		TitleCode: "b_tsahir",
	}
}
