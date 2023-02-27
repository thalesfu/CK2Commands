package aargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳芬贝格LaufenbergBarony struct {
	feud.BaseBarony
}

var BLaufenberg劳芬贝格 feud.Barony = &劳芬贝格LaufenbergBarony{}

func init() {
    f := BLaufenberg劳芬贝格.(*劳芬贝格LaufenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laufenberg",
		TitleName: "劳芬贝格",
		TitleCode: "b_laufenberg",
	}
}
