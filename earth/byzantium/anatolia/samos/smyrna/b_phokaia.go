package smyrna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲莱雅PhokaiaBarony struct {
	feud.BaseBarony
}

var BPhokaia菲莱雅 feud.Barony = &菲莱雅PhokaiaBarony{}

func init() {
	f := BPhokaia菲莱雅.(*菲莱雅PhokaiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phokaia",
		TitleName: "菲莱雅",
		TitleCode: "b_phokaia",
	}
}
