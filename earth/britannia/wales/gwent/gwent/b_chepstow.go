package gwent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切普斯托ChepstowBarony struct {
	feud.BaseBarony
}

var BChepstow切普斯托 feud.Barony = &切普斯托ChepstowBarony{}

func init() {
	f := BChepstow切普斯托.(*切普斯托ChepstowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chepstow",
		TitleName: "切普斯托",
		TitleCode: "b_chepstow",
	}
}
