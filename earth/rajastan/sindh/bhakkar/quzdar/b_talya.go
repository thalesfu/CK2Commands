package quzdar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多俩TalyaBarony struct {
	feud.BaseBarony
}

var BTalya多俩 feud.Barony = &多俩TalyaBarony{}

func init() {
	f := BTalya多俩.(*多俩TalyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talya",
		TitleName: "多俩",
		TitleCode: "b_talya",
	}
}
