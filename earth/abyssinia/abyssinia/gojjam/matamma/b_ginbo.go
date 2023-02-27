package matamma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 津博GinboBarony struct {
	feud.BaseBarony
}

var BGinbo津博 feud.Barony = &津博GinboBarony{}

func init() {
    f := BGinbo津博.(*津博GinboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ginbo",
		TitleName: "津博",
		TitleCode: "b_ginbo",
	}
}
