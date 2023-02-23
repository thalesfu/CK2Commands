package sopron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考普堡KapuvarBarony struct {
	feud.BaseBarony
}

var BKapuvar考普堡 feud.Barony = &考普堡KapuvarBarony{}

func init() {
	f := BKapuvar考普堡.(*考普堡KapuvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapuvar",
		TitleName: "考普堡",
		TitleCode: "b_kapuvar",
	}
}
