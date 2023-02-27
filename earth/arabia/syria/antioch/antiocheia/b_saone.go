package antiocheia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索恩SaoneBarony struct {
	feud.BaseBarony
}

var BSaone索恩 feud.Barony = &索恩SaoneBarony{}

func init() {
    f := BSaone索恩.(*索恩SaoneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saone",
		TitleName: "索恩",
		TitleCode: "b_saone",
	}
}
