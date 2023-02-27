package memel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼达NidaBarony struct {
	feud.BaseBarony
}

var BNida尼达 feud.Barony = &尼达NidaBarony{}

func init() {
    f := BNida尼达.(*尼达NidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nida",
		TitleName: "尼达",
		TitleCode: "b_nida",
	}
}
