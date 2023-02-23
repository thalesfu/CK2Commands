package artois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴波姆BapaumeBarony struct {
	feud.BaseBarony
}

var BBapaume巴波姆 feud.Barony = &巴波姆BapaumeBarony{}

func init() {
	f := BBapaume巴波姆.(*巴波姆BapaumeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bapaume",
		TitleName: "巴波姆",
		TitleCode: "b_bapaume",
	}
}
