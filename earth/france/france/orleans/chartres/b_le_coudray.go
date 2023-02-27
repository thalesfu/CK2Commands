package chartres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒库德赖Le_coudrayBarony struct {
	feud.BaseBarony
}

var BLe_coudray勒库德赖 feud.Barony = &勒库德赖Le_coudrayBarony{}

func init() {
    f := BLe_coudray勒库德赖.(*勒库德赖Le_coudrayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "le_coudray",
		TitleName: "勒库德赖",
		TitleCode: "b_le_coudray",
	}
}
