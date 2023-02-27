package armail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎巴里KambaliBarony struct {
	feud.BaseBarony
}

var BKambali坎巴里 feud.Barony = &坎巴里KambaliBarony{}

func init() {
    f := BKambali坎巴里.(*坎巴里KambaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kambali",
		TitleName: "坎巴里",
		TitleCode: "b_kambali",
	}
}
