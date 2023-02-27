package vijayawada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘布罗尔ChebroluBarony struct {
	feud.BaseBarony
}

var BChebrolu丘布罗尔 feud.Barony = &丘布罗尔ChebroluBarony{}

func init() {
    f := BChebrolu丘布罗尔.(*丘布罗尔ChebroluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chebrolu",
		TitleName: "丘布罗尔",
		TitleCode: "b_chebrolu",
	}
}
