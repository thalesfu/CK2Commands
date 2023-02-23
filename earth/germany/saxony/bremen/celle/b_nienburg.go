package celle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宁堡NienburgBarony struct {
	feud.BaseBarony
}

var BNienburg宁堡 feud.Barony = &宁堡NienburgBarony{}

func init() {
	f := BNienburg宁堡.(*宁堡NienburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nienburg",
		TitleName: "宁堡",
		TitleCode: "b_nienburg",
	}
}
