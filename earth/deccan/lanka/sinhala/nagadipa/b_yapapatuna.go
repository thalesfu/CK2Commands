package nagadipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶波波都那YapapatunaBarony struct {
	feud.BaseBarony
}

var BYapapatuna耶波波都那 feud.Barony = &耶波波都那YapapatunaBarony{}

func init() {
	f := BYapapatuna耶波波都那.(*耶波波都那YapapatunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yapapatuna",
		TitleName: "耶波波都那",
		TitleCode: "b_yapapatuna",
	}
}
