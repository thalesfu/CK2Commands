package more

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔马KalmarBarony struct {
	feud.BaseBarony
}

var BKalmar卡尔马 feud.Barony = &卡尔马KalmarBarony{}

func init() {
	f := BKalmar卡尔马.(*卡尔马KalmarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalmar",
		TitleName: "卡尔马",
		TitleCode: "b_kalmar",
	}
}
