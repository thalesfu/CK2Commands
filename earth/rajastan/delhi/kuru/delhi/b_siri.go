package delhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯利SiriBarony struct {
	feud.BaseBarony
}

var BSiri斯利 feud.Barony = &斯利SiriBarony{}

func init() {
	f := BSiri斯利.(*斯利SiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siri",
		TitleName: "斯利",
		TitleCode: "b_siri",
	}
}
