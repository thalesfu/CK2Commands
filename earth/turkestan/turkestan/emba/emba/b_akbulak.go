package emba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克布拉克AkbulakBarony struct {
	feud.BaseBarony
}

var BAkbulak阿克布拉克 feud.Barony = &阿克布拉克AkbulakBarony{}

func init() {
    f := BAkbulak阿克布拉克.(*阿克布拉克AkbulakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akbulak",
		TitleName: "阿克布拉克",
		TitleCode: "b_akbulak",
	}
}
