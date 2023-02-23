package aswan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉布萨KalabshaBarony struct {
	feud.BaseBarony
}

var BKalabsha卡拉布萨 feud.Barony = &卡拉布萨KalabshaBarony{}

func init() {
	f := BKalabsha卡拉布萨.(*卡拉布萨KalabshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalabsha",
		TitleName: "卡拉布萨",
		TitleCode: "b_kalabsha",
	}
}
