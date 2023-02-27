package wurttemberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 魏布林根WaiblingenBarony struct {
	feud.BaseBarony
}

var BWaiblingen魏布林根 feud.Barony = &魏布林根WaiblingenBarony{}

func init() {
    f := BWaiblingen魏布林根.(*魏布林根WaiblingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waiblingen",
		TitleName: "魏布林根",
		TitleCode: "b_waiblingen",
	}
}
