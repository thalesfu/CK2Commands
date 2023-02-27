package leoben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃彭施泰因EppensteinBarony struct {
	feud.BaseBarony
}

var BEppenstein埃彭施泰因 feud.Barony = &埃彭施泰因EppensteinBarony{}

func init() {
    f := BEppenstein埃彭施泰因.(*埃彭施泰因EppensteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eppenstein",
		TitleName: "埃彭施泰因",
		TitleCode: "b_eppenstein",
	}
}
