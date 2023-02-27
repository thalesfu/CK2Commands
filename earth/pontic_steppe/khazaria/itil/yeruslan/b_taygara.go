package yeruslan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰加拉TaygaraBarony struct {
	feud.BaseBarony
}

var BTaygara泰加拉 feud.Barony = &泰加拉TaygaraBarony{}

func init() {
    f := BTaygara泰加拉.(*泰加拉TaygaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taygara",
		TitleName: "泰加拉",
		TitleCode: "b_taygara",
	}
}
