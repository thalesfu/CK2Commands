package cremona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马内尔比奥ManerbioBarony struct {
	feud.BaseBarony
}

var BManerbio马内尔比奥 feud.Barony = &马内尔比奥ManerbioBarony{}

func init() {
    f := BManerbio马内尔比奥.(*马内尔比奥ManerbioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manerbio",
		TitleName: "马内尔比奥",
		TitleCode: "b_manerbio",
	}
}
