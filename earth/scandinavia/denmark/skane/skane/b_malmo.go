package skane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔默MalmoBarony struct {
	feud.BaseBarony
}

var BMalmo马尔默 feud.Barony = &马尔默MalmoBarony{}

func init() {
	f := BMalmo马尔默.(*马尔默MalmoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malmo",
		TitleName: "马尔默",
		TitleCode: "b_malmo",
	}
}
