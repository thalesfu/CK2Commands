package jersika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿莱内AleneBarony struct {
	feud.BaseBarony
}

var BAlene阿莱内 feud.Barony = &阿莱内AleneBarony{}

func init() {
    f := BAlene阿莱内.(*阿莱内AleneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alene",
		TitleName: "阿莱内",
		TitleCode: "b_alene",
	}
}
