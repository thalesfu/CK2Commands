package thana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 招罗ChaulBarony struct {
	feud.BaseBarony
}

var BChaul招罗 feud.Barony = &招罗ChaulBarony{}

func init() {
	f := BChaul招罗.(*招罗ChaulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaul",
		TitleName: "招罗",
		TitleCode: "b_chaul",
	}
}
