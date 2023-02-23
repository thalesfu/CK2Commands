package khangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 敖尔吉勒OrgilBarony struct {
	feud.BaseBarony
}

var BOrgil敖尔吉勒 feud.Barony = &敖尔吉勒OrgilBarony{}

func init() {
	f := BOrgil敖尔吉勒.(*敖尔吉勒OrgilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orgil",
		TitleName: "敖尔吉勒",
		TitleCode: "b_orgil",
	}
}
