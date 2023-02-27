package el_rif

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿加迪朗AgadireneBarony struct {
	feud.BaseBarony
}

var BAgadirene阿加迪朗 feud.Barony = &阿加迪朗AgadireneBarony{}

func init() {
    f := BAgadirene阿加迪朗.(*阿加迪朗AgadireneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agadirene",
		TitleName: "阿加迪朗",
		TitleCode: "b_agadirene",
	}
}
