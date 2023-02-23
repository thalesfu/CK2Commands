package oxford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷丁ReadingBarony struct {
	feud.BaseBarony
}

var BReading雷丁 feud.Barony = &雷丁ReadingBarony{}

func init() {
	f := BReading雷丁.(*雷丁ReadingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reading",
		TitleName: "雷丁",
		TitleCode: "b_reading",
	}
}
