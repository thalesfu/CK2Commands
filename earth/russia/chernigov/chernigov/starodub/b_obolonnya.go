package starodub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥博隆尼亚ObolonnyaBarony struct {
	feud.BaseBarony
}

var BObolonnya奥博隆尼亚 feud.Barony = &奥博隆尼亚ObolonnyaBarony{}

func init() {
    f := BObolonnya奥博隆尼亚.(*奥博隆尼亚ObolonnyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "obolonnya",
		TitleName: "奥博隆尼亚",
		TitleCode: "b_obolonnya",
	}
}
