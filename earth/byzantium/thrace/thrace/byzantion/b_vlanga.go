package byzantion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌兰戈VlangaBarony struct {
	feud.BaseBarony
}

var BVlanga乌兰戈 feud.Barony = &乌兰戈VlangaBarony{}

func init() {
	f := BVlanga乌兰戈.(*乌兰戈VlangaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vlanga",
		TitleName: "乌兰戈",
		TitleCode: "b_vlanga",
	}
}
