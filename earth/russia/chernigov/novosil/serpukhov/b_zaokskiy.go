package serpukhov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎奥克斯基ZaokskiyBarony struct {
	feud.BaseBarony
}

var BZaokskiy扎奥克斯基 feud.Barony = &扎奥克斯基ZaokskiyBarony{}

func init() {
    f := BZaokskiy扎奥克斯基.(*扎奥克斯基ZaokskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaokskiy",
		TitleName: "扎奥克斯基",
		TitleCode: "b_zaokskiy",
	}
}
