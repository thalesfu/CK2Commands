package murzuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马萨MassahBarony struct {
	feud.BaseBarony
}

var BMassah马萨 feud.Barony = &马萨MassahBarony{}

func init() {
	f := BMassah马萨.(*马萨MassahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "massah",
		TitleName: "马萨",
		TitleCode: "b_massah",
	}
}
