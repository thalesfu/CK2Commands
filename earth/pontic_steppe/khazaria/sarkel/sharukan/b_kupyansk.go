package sharukan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库普扬斯克KupyanskBarony struct {
	feud.BaseBarony
}

var BKupyansk库普扬斯克 feud.Barony = &库普扬斯克KupyanskBarony{}

func init() {
    f := BKupyansk库普扬斯克.(*库普扬斯克KupyanskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kupyansk",
		TitleName: "库普扬斯克",
		TitleCode: "b_kupyansk",
	}
}
