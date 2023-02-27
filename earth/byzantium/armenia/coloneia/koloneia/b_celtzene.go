package koloneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔齐尼CeltzeneBarony struct {
	feud.BaseBarony
}

var BCeltzene凯尔齐尼 feud.Barony = &凯尔齐尼CeltzeneBarony{}

func init() {
    f := BCeltzene凯尔齐尼.(*凯尔齐尼CeltzeneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "celtzene",
		TitleName: "凯尔齐尼",
		TitleCode: "b_celtzene",
	}
}
