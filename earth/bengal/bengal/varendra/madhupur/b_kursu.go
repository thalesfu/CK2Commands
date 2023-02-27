package madhupur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸠须KursuBarony struct {
	feud.BaseBarony
}

var BKursu鸠须 feud.Barony = &鸠须KursuBarony{}

func init() {
    f := BKursu鸠须.(*鸠须KursuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kursu",
		TitleName: "鸠须",
		TitleCode: "b_kursu",
	}
}
