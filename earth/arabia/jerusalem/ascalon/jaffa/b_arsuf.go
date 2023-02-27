package jaffa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔苏夫ArsufBarony struct {
	feud.BaseBarony
}

var BArsuf阿尔苏夫 feud.Barony = &阿尔苏夫ArsufBarony{}

func init() {
    f := BArsuf阿尔苏夫.(*阿尔苏夫ArsufBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arsuf",
		TitleName: "阿尔苏夫",
		TitleCode: "b_arsuf",
	}
}
