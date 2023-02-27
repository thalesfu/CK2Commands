package talas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克丘拉克AkchulakBarony struct {
	feud.BaseBarony
}

var BAkchulak阿克丘拉克 feud.Barony = &阿克丘拉克AkchulakBarony{}

func init() {
    f := BAkchulak阿克丘拉克.(*阿克丘拉克AkchulakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akchulak",
		TitleName: "阿克丘拉克",
		TitleCode: "b_akchulak",
	}
}
