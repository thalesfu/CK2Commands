package tirol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊姆斯特ImstBarony struct {
	feud.BaseBarony
}

var BImst伊姆斯特 feud.Barony = &伊姆斯特ImstBarony{}

func init() {
    f := BImst伊姆斯特.(*伊姆斯特ImstBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "imst",
		TitleName: "伊姆斯特",
		TitleCode: "b_imst",
	}
}
