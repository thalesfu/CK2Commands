package senlis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷皮CrepyBarony struct {
	feud.BaseBarony
}

var BCrepy克雷皮 feud.Barony = &克雷皮CrepyBarony{}

func init() {
	f := BCrepy克雷皮.(*克雷皮CrepyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crepy",
		TitleName: "克雷皮",
		TitleCode: "b_crepy",
	}
}
