package kufa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库法KufaBarony struct {
	feud.BaseBarony
}

var BKufa库法 feud.Barony = &库法KufaBarony{}

func init() {
    f := BKufa库法.(*库法KufaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kufa",
		TitleName: "库法",
		TitleCode: "b_kufa",
	}
}
