package mallorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔盖达AlgaidaBarony struct {
	feud.BaseBarony
}

var BAlgaida阿尔盖达 feud.Barony = &阿尔盖达AlgaidaBarony{}

func init() {
    f := BAlgaida阿尔盖达.(*阿尔盖达AlgaidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "algaida",
		TitleName: "阿尔盖达",
		TitleCode: "b_algaida",
	}
}
