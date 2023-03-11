package cherven

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松夏德卡SaciaskaBarony struct {
	feud.BaseBarony
}

var BSaciaska松夏德卡 feud.Barony = &松夏德卡SaciaskaBarony{}

func init() {
    f := BSaciaska松夏德卡.(*松夏德卡SaciaskaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saciaska",
		TitleName: "松夏德卡",
		TitleCode: "b_saciaska",
	}
}
