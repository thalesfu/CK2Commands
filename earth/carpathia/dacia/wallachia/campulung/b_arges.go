package campulung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔杰什ArgesBarony struct {
	feud.BaseBarony
}

var BArges阿尔杰什 feud.Barony = &阿尔杰什ArgesBarony{}

func init() {
    f := BArges阿尔杰什.(*阿尔杰什ArgesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arges",
		TitleName: "阿尔杰什",
		TitleCode: "b_arges",
	}
}
