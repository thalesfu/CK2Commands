package mecca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麦加MeccaBarony struct {
	feud.BaseBarony
}

var BMecca麦加 feud.Barony = &麦加MeccaBarony{}

func init() {
    f := BMecca麦加.(*麦加MeccaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mecca",
		TitleName: "麦加",
		TitleCode: "b_mecca",
	}
}
