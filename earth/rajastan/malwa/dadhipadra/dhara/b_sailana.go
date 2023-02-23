package dhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨罗那SailanaBarony struct {
	feud.BaseBarony
}

var BSailana萨罗那 feud.Barony = &萨罗那SailanaBarony{}

func init() {
	f := BSailana萨罗那.(*萨罗那SailanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sailana",
		TitleName: "萨罗那",
		TitleCode: "b_sailana",
	}
}
