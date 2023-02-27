package doti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽兰迦KhalangaBarony struct {
	feud.BaseBarony
}

var BKhalanga伽兰迦 feud.Barony = &伽兰迦KhalangaBarony{}

func init() {
    f := BKhalanga伽兰迦.(*伽兰迦KhalangaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khalanga",
		TitleName: "伽兰迦",
		TitleCode: "b_khalanga",
	}
}
