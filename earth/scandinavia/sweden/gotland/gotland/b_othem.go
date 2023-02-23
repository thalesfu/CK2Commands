package gotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌特海姆OthemBarony struct {
	feud.BaseBarony
}

var BOthem乌特海姆 feud.Barony = &乌特海姆OthemBarony{}

func init() {
	f := BOthem乌特海姆.(*乌特海姆OthemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "othem",
		TitleName: "乌特海姆",
		TitleCode: "b_othem",
	}
}
