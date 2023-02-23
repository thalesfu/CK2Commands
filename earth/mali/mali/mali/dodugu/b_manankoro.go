package dodugu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马南科罗ManankoroBarony struct {
	feud.BaseBarony
}

var BManankoro马南科罗 feud.Barony = &马南科罗ManankoroBarony{}

func init() {
	f := BManankoro马南科罗.(*马南科罗ManankoroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manankoro",
		TitleName: "马南科罗",
		TitleCode: "b_manankoro",
	}
}
