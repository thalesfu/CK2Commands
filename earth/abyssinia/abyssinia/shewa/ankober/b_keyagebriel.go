package ankober

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克亚吉比利KeyagebrielBarony struct {
	feud.BaseBarony
}

var BKeyagebriel克亚吉比利 feud.Barony = &克亚吉比利KeyagebrielBarony{}

func init() {
	f := BKeyagebriel克亚吉比利.(*克亚吉比利KeyagebrielBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keyagebriel",
		TitleName: "克亚吉比利",
		TitleCode: "b_keyagebriel",
	}
}
