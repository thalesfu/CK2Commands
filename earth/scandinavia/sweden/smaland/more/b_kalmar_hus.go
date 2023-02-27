package more

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔马胡斯Kalmar_husBarony struct {
	feud.BaseBarony
}

var BKalmar_hus卡尔马胡斯 feud.Barony = &卡尔马胡斯Kalmar_husBarony{}

func init() {
    f := BKalmar_hus卡尔马胡斯.(*卡尔马胡斯Kalmar_husBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalmar_hus",
		TitleName: "卡尔马胡斯",
		TitleCode: "b_kalmar_hus",
	}
}
