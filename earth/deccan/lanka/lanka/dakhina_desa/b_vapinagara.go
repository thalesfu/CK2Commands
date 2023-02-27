package dakhina_desa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆比那揭罗VapinagaraBarony struct {
	feud.BaseBarony
}

var BVapinagara婆比那揭罗 feud.Barony = &婆比那揭罗VapinagaraBarony{}

func init() {
    f := BVapinagara婆比那揭罗.(*婆比那揭罗VapinagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vapinagara",
		TitleName: "婆比那揭罗",
		TitleCode: "b_vapinagara",
	}
}
