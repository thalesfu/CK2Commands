package edessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨姆萨特SamsatBarony struct {
	feud.BaseBarony
}

var BSamsat萨姆萨特 feud.Barony = &萨姆萨特SamsatBarony{}

func init() {
	f := BSamsat萨姆萨特.(*萨姆萨特SamsatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samsat",
		TitleName: "萨姆萨特",
		TitleCode: "b_samsat",
	}
}
