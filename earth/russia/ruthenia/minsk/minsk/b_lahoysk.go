package minsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛戈伊斯克LahoyskBarony struct {
	feud.BaseBarony
}

var BLahoysk洛戈伊斯克 feud.Barony = &洛戈伊斯克LahoyskBarony{}

func init() {
    f := BLahoysk洛戈伊斯克.(*洛戈伊斯克LahoyskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lahoysk",
		TitleName: "洛戈伊斯克",
		TitleCode: "b_lahoysk",
	}
}
