package karbala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙弥亚ShamiyahBarony struct {
	feud.BaseBarony
}

var BShamiyah沙弥亚 feud.Barony = &沙弥亚ShamiyahBarony{}

func init() {
    f := BShamiyah沙弥亚.(*沙弥亚ShamiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shamiyah",
		TitleName: "沙弥亚",
		TitleCode: "b_shamiyah",
	}
}
