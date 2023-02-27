package shahrazur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯克拉巴德EskelabadBarony struct {
	feud.BaseBarony
}

var BEskelabad埃斯克拉巴德 feud.Barony = &埃斯克拉巴德EskelabadBarony{}

func init() {
    f := BEskelabad埃斯克拉巴德.(*埃斯克拉巴德EskelabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eskelabad",
		TitleName: "埃斯克拉巴德",
		TitleCode: "b_eskelabad",
	}
}
