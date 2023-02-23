package shigatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 三竹节SamzhubzeBarony struct {
	feud.BaseBarony
}

var BSamzhubze三竹节 feud.Barony = &三竹节SamzhubzeBarony{}

func init() {
	f := BSamzhubze三竹节.(*三竹节SamzhubzeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samzhubze",
		TitleName: "三竹节",
		TitleCode: "b_samzhubze",
	}
}
