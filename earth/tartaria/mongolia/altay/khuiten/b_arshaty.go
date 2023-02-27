package khuiten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔沙特ArshatyBarony struct {
	feud.BaseBarony
}

var BArshaty阿尔沙特 feud.Barony = &阿尔沙特ArshatyBarony{}

func init() {
    f := BArshaty阿尔沙特.(*阿尔沙特ArshatyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arshaty",
		TitleName: "阿尔沙特",
		TitleCode: "b_arshaty",
	}
}
