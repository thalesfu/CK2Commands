package tabriz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎哈克ZahhakBarony struct {
	feud.BaseBarony
}

var BZahhak扎哈克 feud.Barony = &扎哈克ZahhakBarony{}

func init() {
    f := BZahhak扎哈克.(*扎哈克ZahhakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zahhak",
		TitleName: "扎哈克",
		TitleCode: "b_zahhak",
	}
}
