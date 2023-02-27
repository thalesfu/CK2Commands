package leoben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加仑施泰因GallensteinBarony struct {
	feud.BaseBarony
}

var BGallenstein加仑施泰因 feud.Barony = &加仑施泰因GallensteinBarony{}

func init() {
    f := BGallenstein加仑施泰因.(*加仑施泰因GallensteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gallenstein",
		TitleName: "加仑施泰因",
		TitleCode: "b_gallenstein",
	}
}
