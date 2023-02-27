package kaneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎迪亚KandiaBarony struct {
	feud.BaseBarony
}

var BKandia坎迪亚 feud.Barony = &坎迪亚KandiaBarony{}

func init() {
    f := BKandia坎迪亚.(*坎迪亚KandiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kandia",
		TitleName: "坎迪亚",
		TitleCode: "b_kandia",
	}
}
