package galich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加利奇GalichBarony struct {
	feud.BaseBarony
}

var BGalich加利奇 feud.Barony = &加利奇GalichBarony{}

func init() {
    f := BGalich加利奇.(*加利奇GalichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galich",
		TitleName: "加利奇",
		TitleCode: "b_galich",
	}
}
