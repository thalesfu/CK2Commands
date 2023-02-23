package kalat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琴加ChengaBarony struct {
	feud.BaseBarony
}

var BChenga琴加 feud.Barony = &琴加ChengaBarony{}

func init() {
	f := BChenga琴加.(*琴加ChengaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chenga",
		TitleName: "琴加",
		TitleCode: "b_chenga",
	}
}
