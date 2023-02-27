package maragha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基莱欣Kileh_shinBarony struct {
	feud.BaseBarony
}

var BKileh_shin基莱欣 feud.Barony = &基莱欣Kileh_shinBarony{}

func init() {
    f := BKileh_shin基莱欣.(*基莱欣Kileh_shinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kileh_shin",
		TitleName: "基莱欣",
		TitleCode: "b_kileh_shin",
	}
}
