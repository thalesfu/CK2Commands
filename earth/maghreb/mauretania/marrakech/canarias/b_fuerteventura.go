package canarias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富埃特文图拉FuerteventuraBarony struct {
	feud.BaseBarony
}

var BFuerteventura富埃特文图拉 feud.Barony = &富埃特文图拉FuerteventuraBarony{}

func init() {
    f := BFuerteventura富埃特文图拉.(*富埃特文图拉FuerteventuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuerteventura",
		TitleName: "富埃特文图拉",
		TitleCode: "b_fuerteventura",
	}
}
