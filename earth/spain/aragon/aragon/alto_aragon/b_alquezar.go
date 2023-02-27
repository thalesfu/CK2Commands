package alto_aragon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔克萨尔AlquezarBarony struct {
	feud.BaseBarony
}

var BAlquezar阿尔克萨尔 feud.Barony = &阿尔克萨尔AlquezarBarony{}

func init() {
    f := BAlquezar阿尔克萨尔.(*阿尔克萨尔AlquezarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alquezar",
		TitleName: "阿尔克萨尔",
		TitleCode: "b_alquezar",
	}
}
