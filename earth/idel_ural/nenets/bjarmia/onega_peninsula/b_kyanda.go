package onega_peninsula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基扬达KyandaBarony struct {
	feud.BaseBarony
}

var BKyanda基扬达 feud.Barony = &基扬达KyandaBarony{}

func init() {
    f := BKyanda基扬达.(*基扬达KyandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyanda",
		TitleName: "基扬达",
		TitleCode: "b_kyanda",
	}
}
