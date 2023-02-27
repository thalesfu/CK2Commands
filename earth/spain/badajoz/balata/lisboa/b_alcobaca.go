package lisboa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔科巴萨AlcobacaBarony struct {
	feud.BaseBarony
}

var BAlcobaca阿尔科巴萨 feud.Barony = &阿尔科巴萨AlcobacaBarony{}

func init() {
    f := BAlcobaca阿尔科巴萨.(*阿尔科巴萨AlcobacaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcobaca",
		TitleName: "阿尔科巴萨",
		TitleCode: "b_alcobaca",
	}
}
