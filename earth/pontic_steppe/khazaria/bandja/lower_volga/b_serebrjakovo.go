package lower_volga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢列布里亚科沃SerebrjakovoBarony struct {
	feud.BaseBarony
}

var BSerebrjakovo谢列布里亚科沃 feud.Barony = &谢列布里亚科沃SerebrjakovoBarony{}

func init() {
    f := BSerebrjakovo谢列布里亚科沃.(*谢列布里亚科沃SerebrjakovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serebrjakovo",
		TitleName: "谢列布里亚科沃",
		TitleCode: "b_serebrjakovo",
	}
}
