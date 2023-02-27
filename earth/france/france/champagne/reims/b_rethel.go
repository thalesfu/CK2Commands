package reims

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒泰勒RethelBarony struct {
	feud.BaseBarony
}

var BRethel勒泰勒 feud.Barony = &勒泰勒RethelBarony{}

func init() {
    f := BRethel勒泰勒.(*勒泰勒RethelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rethel",
		TitleName: "勒泰勒",
		TitleCode: "b_rethel",
	}
}
