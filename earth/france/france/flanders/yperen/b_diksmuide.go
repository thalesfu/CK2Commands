package yperen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪克斯迈德DiksmuideBarony struct {
	feud.BaseBarony
}

var BDiksmuide迪克斯迈德 feud.Barony = &迪克斯迈德DiksmuideBarony{}

func init() {
	f := BDiksmuide迪克斯迈德.(*迪克斯迈德DiksmuideBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diksmuide",
		TitleName: "迪克斯迈德",
		TitleCode: "b_diksmuide",
	}
}
