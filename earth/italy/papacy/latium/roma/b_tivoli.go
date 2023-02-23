package roma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂沃利TivoliBarony struct {
	feud.BaseBarony
}

var BTivoli蒂沃利 feud.Barony = &蒂沃利TivoliBarony{}

func init() {
	f := BTivoli蒂沃利.(*蒂沃利TivoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tivoli",
		TitleName: "蒂沃利",
		TitleCode: "b_tivoli",
	}
}
