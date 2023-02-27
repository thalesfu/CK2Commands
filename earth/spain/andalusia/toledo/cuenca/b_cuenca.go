package cuenca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昆卡CuencaBarony struct {
	feud.BaseBarony
}

var BCuenca昆卡 feud.Barony = &昆卡CuencaBarony{}

func init() {
    f := BCuenca昆卡.(*昆卡CuencaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cuenca",
		TitleName: "昆卡",
		TitleCode: "b_cuenca",
	}
}
