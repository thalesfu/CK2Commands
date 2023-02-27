package ishim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶西尔EsilBarony struct {
	feud.BaseBarony
}

var BEsil叶西尔 feud.Barony = &叶西尔EsilBarony{}

func init() {
    f := BEsil叶西尔.(*叶西尔EsilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esil",
		TitleName: "叶西尔",
		TitleCode: "b_esil",
	}
}
