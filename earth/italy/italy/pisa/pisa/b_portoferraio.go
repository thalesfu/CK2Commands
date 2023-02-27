package pisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费拉约港PortoferraioBarony struct {
	feud.BaseBarony
}

var BPortoferraio费拉约港 feud.Barony = &费拉约港PortoferraioBarony{}

func init() {
    f := BPortoferraio费拉约港.(*费拉约港PortoferraioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portoferraio",
		TitleName: "费拉约港",
		TitleCode: "b_portoferraio",
	}
}
