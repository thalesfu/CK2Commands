package tirgoviste

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普楚尔布蒂Tabla_butiiBarony struct {
	feud.BaseBarony
}

var BTabla_butii普楚尔布蒂 feud.Barony = &普楚尔布蒂Tabla_butiiBarony{}

func init() {
    f := BTabla_butii普楚尔布蒂.(*普楚尔布蒂Tabla_butiiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tabla_butii",
		TitleName: "普楚尔布蒂",
		TitleCode: "b_tabla_butii",
	}
}
