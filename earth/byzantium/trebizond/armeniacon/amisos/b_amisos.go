package amisos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿米索斯AmisosBarony struct {
	feud.BaseBarony
}

var BAmisos阿米索斯 feud.Barony = &阿米索斯AmisosBarony{}

func init() {
	f := BAmisos阿米索斯.(*阿米索斯AmisosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amisos",
		TitleName: "阿米索斯",
		TitleCode: "b_amisos",
	}
}
