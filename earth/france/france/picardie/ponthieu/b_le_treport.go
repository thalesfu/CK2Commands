package ponthieu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒特雷波尔Le_treportBarony struct {
	feud.BaseBarony
}

var BLe_treport勒特雷波尔 feud.Barony = &勒特雷波尔Le_treportBarony{}

func init() {
    f := BLe_treport勒特雷波尔.(*勒特雷波尔Le_treportBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "le_treport",
		TitleName: "勒特雷波尔",
		TitleCode: "b_le_treport",
	}
}
