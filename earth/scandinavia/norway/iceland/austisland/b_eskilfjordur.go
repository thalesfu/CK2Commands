package austisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯基尔菲厄泽EskilfjordurBarony struct {
	feud.BaseBarony
}

var BEskilfjordur埃斯基尔菲厄泽 feud.Barony = &埃斯基尔菲厄泽EskilfjordurBarony{}

func init() {
    f := BEskilfjordur埃斯基尔菲厄泽.(*埃斯基尔菲厄泽EskilfjordurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eskilfjordur",
		TitleName: "埃斯基尔菲厄泽",
		TitleCode: "b_eskilfjordur",
	}
}
