package wana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽蓝菩梨KalamboliBarony struct {
	feud.BaseBarony
}

var BKalamboli伽蓝菩梨 feud.Barony = &伽蓝菩梨KalamboliBarony{}

func init() {
	f := BKalamboli伽蓝菩梨.(*伽蓝菩梨KalamboliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalamboli",
		TitleName: "伽蓝菩梨",
		TitleCode: "b_kalamboli",
	}
}
