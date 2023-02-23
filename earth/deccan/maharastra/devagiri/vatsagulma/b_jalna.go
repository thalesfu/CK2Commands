package vatsagulma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇罗那JalnaBarony struct {
	feud.BaseBarony
}

var BJalna阇罗那 feud.Barony = &阇罗那JalnaBarony{}

func init() {
	f := BJalna阇罗那.(*阇罗那JalnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jalna",
		TitleName: "阇罗那",
		TitleCode: "b_jalna",
	}
}
