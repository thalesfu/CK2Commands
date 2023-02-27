package bizerte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈茨阿NetzaBarony struct {
	feud.BaseBarony
}

var BNetza奈茨阿 feud.Barony = &奈茨阿NetzaBarony{}

func init() {
    f := BNetza奈茨阿.(*奈茨阿NetzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "netza",
		TitleName: "奈茨阿",
		TitleCode: "b_netza",
	}
}
