package sarpa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊基布鲁尔IkburulBarony struct {
	feud.BaseBarony
}

var BIkburul伊基布鲁尔 feud.Barony = &伊基布鲁尔IkburulBarony{}

func init() {
    f := BIkburul伊基布鲁尔.(*伊基布鲁尔IkburulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikburul",
		TitleName: "伊基布鲁尔",
		TitleCode: "b_ikburul",
	}
}
