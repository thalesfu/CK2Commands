package dal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔斯波TossboBarony struct {
	feud.BaseBarony
}

var BTossbo塔斯波 feud.Barony = &塔斯波TossboBarony{}

func init() {
    f := BTossbo塔斯波.(*塔斯波TossboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tossbo",
		TitleName: "塔斯波",
		TitleCode: "b_tossbo",
	}
}
