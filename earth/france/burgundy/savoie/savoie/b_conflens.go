package savoie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔福朗ConflensBarony struct {
	feud.BaseBarony
}

var BConflens孔福朗 feud.Barony = &孔福朗ConflensBarony{}

func init() {
    f := BConflens孔福朗.(*孔福朗ConflensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "conflens",
		TitleName: "孔福朗",
		TitleCode: "b_conflens",
	}
}
