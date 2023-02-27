package sakmara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨克马拉SakmaraBarony struct {
	feud.BaseBarony
}

var BSakmara萨克马拉 feud.Barony = &萨克马拉SakmaraBarony{}

func init() {
    f := BSakmara萨克马拉.(*萨克马拉SakmaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sakmara",
		TitleName: "萨克马拉",
		TitleCode: "b_sakmara",
	}
}
