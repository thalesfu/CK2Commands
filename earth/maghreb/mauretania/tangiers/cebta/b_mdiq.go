package cebta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈迪格MdiqBarony struct {
	feud.BaseBarony
}

var BMdiq迈迪格 feud.Barony = &迈迪格MdiqBarony{}

func init() {
	f := BMdiq迈迪格.(*迈迪格MdiqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mdiq",
		TitleName: "迈迪格",
		TitleCode: "b_mdiq",
	}
}
