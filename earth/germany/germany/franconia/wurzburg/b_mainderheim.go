package wurzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明德尔海姆MainderheimBarony struct {
	feud.BaseBarony
}

var BMainderheim明德尔海姆 feud.Barony = &明德尔海姆MainderheimBarony{}

func init() {
	f := BMainderheim明德尔海姆.(*明德尔海姆MainderheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mainderheim",
		TitleName: "明德尔海姆",
		TitleCode: "b_mainderheim",
	}
}
