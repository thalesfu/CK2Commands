package deir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提克里特TakritBarony struct {
	feud.BaseBarony
}

var BTakrit提克里特 feud.Barony = &提克里特TakritBarony{}

func init() {
    f := BTakrit提克里特.(*提克里特TakritBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "takrit",
		TitleName: "提克里特",
		TitleCode: "b_takrit",
	}
}
