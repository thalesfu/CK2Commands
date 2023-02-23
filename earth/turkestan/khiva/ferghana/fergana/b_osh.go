package fergana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 我失OshBarony struct {
	feud.BaseBarony
}

var BOsh我失 feud.Barony = &我失OshBarony{}

func init() {
	f := BOsh我失.(*我失OshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osh",
		TitleName: "我失",
		TitleCode: "b_osh",
	}
}
