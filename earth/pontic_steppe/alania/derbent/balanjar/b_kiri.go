package balanjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基里KiriBarony struct {
	feud.BaseBarony
}

var BKiri基里 feud.Barony = &基里KiriBarony{}

func init() {
    f := BKiri基里.(*基里KiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiri",
		TitleName: "基里",
		TitleCode: "b_kiri",
	}
}
