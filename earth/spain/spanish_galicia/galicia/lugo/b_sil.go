package lugo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西尔SilBarony struct {
	feud.BaseBarony
}

var BSil西尔 feud.Barony = &西尔SilBarony{}

func init() {
    f := BSil西尔.(*西尔SilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sil",
		TitleName: "西尔",
		TitleCode: "b_sil",
	}
}
