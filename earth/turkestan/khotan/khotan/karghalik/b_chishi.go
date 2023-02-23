package karghalik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赤石ChishiBarony struct {
	feud.BaseBarony
}

var BChishi赤石 feud.Barony = &赤石ChishiBarony{}

func init() {
	f := BChishi赤石.(*赤石ChishiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chishi",
		TitleName: "赤石",
		TitleCode: "b_chishi",
	}
}
