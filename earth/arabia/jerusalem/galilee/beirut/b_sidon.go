package beirut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西顿SidonBarony struct {
	feud.BaseBarony
}

var BSidon西顿 feud.Barony = &西顿SidonBarony{}

func init() {
    f := BSidon西顿.(*西顿SidonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidon",
		TitleName: "西顿",
		TitleCode: "b_sidon",
	}
}
