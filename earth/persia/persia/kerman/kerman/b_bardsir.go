package kerman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔德斯尔BardsirBarony struct {
	feud.BaseBarony
}

var BBardsir巴尔德斯尔 feud.Barony = &巴尔德斯尔BardsirBarony{}

func init() {
	f := BBardsir巴尔德斯尔.(*巴尔德斯尔BardsirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bardsir",
		TitleName: "巴尔德斯尔",
		TitleCode: "b_bardsir",
	}
}
