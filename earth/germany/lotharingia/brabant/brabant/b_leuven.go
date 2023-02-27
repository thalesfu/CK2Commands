package brabant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒芬LeuvenBarony struct {
	feud.BaseBarony
}

var BLeuven勒芬 feud.Barony = &勒芬LeuvenBarony{}

func init() {
    f := BLeuven勒芬.(*勒芬LeuvenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leuven",
		TitleName: "勒芬",
		TitleCode: "b_leuven",
	}
}
