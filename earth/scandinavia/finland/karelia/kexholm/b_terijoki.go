package kexholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰里约基TerijokiBarony struct {
	feud.BaseBarony
}

var BTerijoki泰里约基 feud.Barony = &泰里约基TerijokiBarony{}

func init() {
    f := BTerijoki泰里约基.(*泰里约基TerijokiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terijoki",
		TitleName: "泰里约基",
		TitleCode: "b_terijoki",
	}
}
