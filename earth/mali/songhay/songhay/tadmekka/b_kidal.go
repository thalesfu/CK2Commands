package tadmekka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基达勒KidalBarony struct {
	feud.BaseBarony
}

var BKidal基达勒 feud.Barony = &基达勒KidalBarony{}

func init() {
	f := BKidal基达勒.(*基达勒KidalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kidal",
		TitleName: "基达勒",
		TitleCode: "b_kidal",
	}
}
