package ajayameru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那丽利NareliBarony struct {
	feud.BaseBarony
}

var BNareli那丽利 feud.Barony = &那丽利NareliBarony{}

func init() {
	f := BNareli那丽利.(*那丽利NareliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nareli",
		TitleName: "那丽利",
		TitleCode: "b_nareli",
	}
}
