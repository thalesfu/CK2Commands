package novgorod_seversky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科诺托普KonotopBarony struct {
	feud.BaseBarony
}

var BKonotop科诺托普 feud.Barony = &科诺托普KonotopBarony{}

func init() {
    f := BKonotop科诺托普.(*科诺托普KonotopBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konotop",
		TitleName: "科诺托普",
		TitleCode: "b_konotop",
	}
}
