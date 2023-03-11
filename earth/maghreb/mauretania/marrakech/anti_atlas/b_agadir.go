package anti_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿加迪尔AgadirBarony struct {
	feud.BaseBarony
}

var BAgadir阿加迪尔 feud.Barony = &阿加迪尔AgadirBarony{}

func init() {
    f := BAgadir阿加迪尔.(*阿加迪尔AgadirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agadir",
		TitleName: "阿加迪尔",
		TitleCode: "b_agadir",
	}
}
