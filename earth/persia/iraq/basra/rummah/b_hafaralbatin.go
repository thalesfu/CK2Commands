package rummah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈夫阿巴廷HafaralbatinBarony struct {
	feud.BaseBarony
}

var BHafaralbatin哈夫阿巴廷 feud.Barony = &哈夫阿巴廷HafaralbatinBarony{}

func init() {
	f := BHafaralbatin哈夫阿巴廷.(*哈夫阿巴廷HafaralbatinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hafaralbatin",
		TitleName: "哈夫阿巴廷",
		TitleCode: "b_hafaralbatin",
	}
}
