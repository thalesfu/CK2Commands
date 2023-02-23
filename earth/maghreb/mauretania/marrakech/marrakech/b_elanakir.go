package marrakech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴基ElanakirBarony struct {
	feud.BaseBarony
}

var BElanakir阿巴基 feud.Barony = &阿巴基ElanakirBarony{}

func init() {
	f := BElanakir阿巴基.(*阿巴基ElanakirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elanakir",
		TitleName: "阿巴基",
		TitleCode: "b_elanakir",
	}
}
