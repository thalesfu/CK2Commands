package ikonion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安姆布拉达AmbladaBarony struct {
	feud.BaseBarony
}

var BAmblada安姆布拉达 feud.Barony = &安姆布拉达AmbladaBarony{}

func init() {
    f := BAmblada安姆布拉达.(*安姆布拉达AmbladaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amblada",
		TitleName: "安姆布拉达",
		TitleCode: "b_amblada",
	}
}
