package faereyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍夫HovBarony struct {
	feud.BaseBarony
}

var BHov霍夫 feud.Barony = &霍夫HovBarony{}

func init() {
	f := BHov霍夫.(*霍夫HovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hov",
		TitleName: "霍夫",
		TitleCode: "b_hov",
	}
}
