package podlasie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克伦基KrynkiBarony struct {
	feud.BaseBarony
}

var BKrynki克伦基 feud.Barony = &克伦基KrynkiBarony{}

func init() {
    f := BKrynki克伦基.(*克伦基KrynkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krynki",
		TitleName: "克伦基",
		TitleCode: "b_krynki",
	}
}
