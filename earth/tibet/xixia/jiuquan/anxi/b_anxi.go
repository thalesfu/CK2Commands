package anxi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安西AnxiBarony struct {
	feud.BaseBarony
}

var BAnxi安西 feud.Barony = &安西AnxiBarony{}

func init() {
	f := BAnxi安西.(*安西AnxiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anxi",
		TitleName: "安西",
		TitleCode: "b_anxi",
	}
}
