package bauchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 包奇BauchiBarony struct {
	feud.BaseBarony
}

var BBauchi包奇 feud.Barony = &包奇BauchiBarony{}

func init() {
    f := BBauchi包奇.(*包奇BauchiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bauchi",
		TitleName: "包奇",
		TitleCode: "b_bauchi",
	}
}
