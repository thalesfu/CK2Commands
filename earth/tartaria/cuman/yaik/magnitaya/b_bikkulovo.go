package magnitaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比库洛沃BikkulovoBarony struct {
	feud.BaseBarony
}

var BBikkulovo比库洛沃 feud.Barony = &比库洛沃BikkulovoBarony{}

func init() {
	f := BBikkulovo比库洛沃.(*比库洛沃BikkulovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bikkulovo",
		TitleName: "比库洛沃",
		TitleCode: "b_bikkulovo",
	}
}
