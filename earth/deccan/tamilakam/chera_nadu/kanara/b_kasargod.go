package kanara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽萨罗KasargodBarony struct {
	feud.BaseBarony
}

var BKasargod伽萨罗 feud.Barony = &伽萨罗KasargodBarony{}

func init() {
    f := BKasargod伽萨罗.(*伽萨罗KasargodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kasargod",
		TitleName: "伽萨罗",
		TitleCode: "b_kasargod",
	}
}
