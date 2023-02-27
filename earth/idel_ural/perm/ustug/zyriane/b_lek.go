package zyriane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列克LekBarony struct {
	feud.BaseBarony
}

var BLek列克 feud.Barony = &列克LekBarony{}

func init() {
    f := BLek列克.(*列克LekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lek",
		TitleName: "列克",
		TitleCode: "b_lek",
	}
}
