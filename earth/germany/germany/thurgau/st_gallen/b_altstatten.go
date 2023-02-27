package st_gallen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔特施泰滕AltstattenBarony struct {
	feud.BaseBarony
}

var BAltstatten阿尔特施泰滕 feud.Barony = &阿尔特施泰滕AltstattenBarony{}

func init() {
    f := BAltstatten阿尔特施泰滕.(*阿尔特施泰滕AltstattenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altstatten",
		TitleName: "阿尔特施泰滕",
		TitleCode: "b_altstatten",
	}
}
