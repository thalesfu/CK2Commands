package thana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴森BasseinBarony struct {
	feud.BaseBarony
}

var BBassein巴森 feud.Barony = &巴森BasseinBarony{}

func init() {
	f := BBassein巴森.(*巴森BasseinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bassein",
		TitleName: "巴森",
		TitleCode: "b_bassein",
	}
}
