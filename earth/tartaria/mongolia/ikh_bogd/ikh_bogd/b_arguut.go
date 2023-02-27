package ikh_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔古特ArguutBarony struct {
	feud.BaseBarony
}

var BArguut阿尔古特 feud.Barony = &阿尔古特ArguutBarony{}

func init() {
    f := BArguut阿尔古特.(*阿尔古特ArguutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arguut",
		TitleName: "阿尔古特",
		TitleCode: "b_arguut",
	}
}
