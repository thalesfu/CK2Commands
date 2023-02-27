package huelva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿亚蒙特AyamonteBarony struct {
	feud.BaseBarony
}

var BAyamonte阿亚蒙特 feud.Barony = &阿亚蒙特AyamonteBarony{}

func init() {
    f := BAyamonte阿亚蒙特.(*阿亚蒙特AyamonteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ayamonte",
		TitleName: "阿亚蒙特",
		TitleCode: "b_ayamonte",
	}
}
