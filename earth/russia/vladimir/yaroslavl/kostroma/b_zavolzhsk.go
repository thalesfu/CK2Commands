package kostroma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎沃尔日斯克ZavolzhskBarony struct {
	feud.BaseBarony
}

var BZavolzhsk扎沃尔日斯克 feud.Barony = &扎沃尔日斯克ZavolzhskBarony{}

func init() {
	f := BZavolzhsk扎沃尔日斯克.(*扎沃尔日斯克ZavolzhskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zavolzhsk",
		TitleName: "扎沃尔日斯克",
		TitleCode: "b_zavolzhsk",
	}
}
