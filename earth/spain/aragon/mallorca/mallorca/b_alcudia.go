package mallorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔库迪亚AlcudiaBarony struct {
	feud.BaseBarony
}

var BAlcudia阿尔库迪亚 feud.Barony = &阿尔库迪亚AlcudiaBarony{}

func init() {
    f := BAlcudia阿尔库迪亚.(*阿尔库迪亚AlcudiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcudia",
		TitleName: "阿尔库迪亚",
		TitleCode: "b_alcudia",
	}
}
