package eilat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨皮尔SapirBarony struct {
	feud.BaseBarony
}

var BSapir萨皮尔 feud.Barony = &萨皮尔SapirBarony{}

func init() {
	f := BSapir萨皮尔.(*萨皮尔SapirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sapir",
		TitleName: "萨皮尔",
		TitleCode: "b_sapir",
	}
}
