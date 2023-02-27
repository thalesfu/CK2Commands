package bukhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉日杜万GizhduvanBarony struct {
	feud.BaseBarony
}

var BGizhduvan吉日杜万 feud.Barony = &吉日杜万GizhduvanBarony{}

func init() {
    f := BGizhduvan吉日杜万.(*吉日杜万GizhduvanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gizhduvan",
		TitleName: "吉日杜万",
		TitleCode: "b_gizhduvan",
	}
}
