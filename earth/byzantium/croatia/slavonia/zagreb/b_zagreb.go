package zagreb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎格雷布ZagrebBarony struct {
	feud.BaseBarony
}

var BZagreb扎格雷布 feud.Barony = &扎格雷布ZagrebBarony{}

func init() {
    f := BZagreb扎格雷布.(*扎格雷布ZagrebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zagreb",
		TitleName: "扎格雷布",
		TitleCode: "b_zagreb",
	}
}
