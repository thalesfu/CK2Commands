package mzab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞卜塞卜SebsebBarony struct {
	feud.BaseBarony
}

var BSebseb塞卜塞卜 feud.Barony = &塞卜塞卜SebsebBarony{}

func init() {
	f := BSebseb塞卜塞卜.(*塞卜塞卜SebsebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sebseb",
		TitleName: "塞卜塞卜",
		TitleCode: "b_sebseb",
	}
}
