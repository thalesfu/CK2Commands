package annaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布乌贾马BoudjamaBarony struct {
	feud.BaseBarony
}

var BBoudjama布乌贾马 feud.Barony = &布乌贾马BoudjamaBarony{}

func init() {
	f := BBoudjama布乌贾马.(*布乌贾马BoudjamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boudjama",
		TitleName: "布乌贾马",
		TitleCode: "b_boudjama",
	}
}
