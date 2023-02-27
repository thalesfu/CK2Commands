package aj_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔班特斯GurvantesBarony struct {
	feud.BaseBarony
}

var BGurvantes古尔班特斯 feud.Barony = &古尔班特斯GurvantesBarony{}

func init() {
    f := BGurvantes古尔班特斯.(*古尔班特斯GurvantesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurvantes",
		TitleName: "古尔班特斯",
		TitleCode: "b_gurvantes",
	}
}
