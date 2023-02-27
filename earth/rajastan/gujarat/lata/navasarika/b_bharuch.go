package navasarika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋禄羯呫婆BharuchBarony struct {
	feud.BaseBarony
}

var BBharuch跋禄羯呫婆 feud.Barony = &跋禄羯呫婆BharuchBarony{}

func init() {
    f := BBharuch跋禄羯呫婆.(*跋禄羯呫婆BharuchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bharuch",
		TitleName: "跋禄羯呫婆",
		TitleCode: "b_bharuch",
	}
}
