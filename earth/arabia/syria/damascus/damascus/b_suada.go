package damascus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏瓦达SuadaBarony struct {
	feud.BaseBarony
}

var BSuada苏瓦达 feud.Barony = &苏瓦达SuadaBarony{}

func init() {
	f := BSuada苏瓦达.(*苏瓦达SuadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suada",
		TitleName: "苏瓦达",
		TitleCode: "b_suada",
	}
}
