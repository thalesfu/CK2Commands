package tobysh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏拉Sula_tobyshBarony struct {
	feud.BaseBarony
}

var BSula_tobysh苏拉 feud.Barony = &苏拉Sula_tobyshBarony{}

func init() {
    f := BSula_tobysh苏拉.(*苏拉Sula_tobyshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sula_tobysh",
		TitleName: "苏拉",
		TitleCode: "b_sula_tobysh",
	}
}
