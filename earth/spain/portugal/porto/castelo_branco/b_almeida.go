package castelo_branco

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔梅达AlmeidaBarony struct {
	feud.BaseBarony
}

var BAlmeida阿尔梅达 feud.Barony = &阿尔梅达AlmeidaBarony{}

func init() {
    f := BAlmeida阿尔梅达.(*阿尔梅达AlmeidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almeida",
		TitleName: "阿尔梅达",
		TitleCode: "b_almeida",
	}
}
