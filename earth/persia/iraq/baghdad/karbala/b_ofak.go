package karbala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥法可OfakBarony struct {
	feud.BaseBarony
}

var BOfak奥法可 feud.Barony = &奥法可OfakBarony{}

func init() {
	f := BOfak奥法可.(*奥法可OfakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ofak",
		TitleName: "奥法可",
		TitleCode: "b_ofak",
	}
}
