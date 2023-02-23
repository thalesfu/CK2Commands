package belgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切达埃斯特CedaestiBarony struct {
	feud.BaseBarony
}

var BCedaesti切达埃斯特 feud.Barony = &切达埃斯特CedaestiBarony{}

func init() {
	f := BCedaesti切达埃斯特.(*切达埃斯特CedaestiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cedaesti",
		TitleName: "切达埃斯特",
		TitleCode: "b_cedaesti",
	}
}
