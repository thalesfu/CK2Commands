package begemder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝特约哈尼斯Bete_yohanisBarony struct {
	feud.BaseBarony
}

var BBete_yohanis贝特约哈尼斯 feud.Barony = &贝特约哈尼斯Bete_yohanisBarony{}

func init() {
    f := BBete_yohanis贝特约哈尼斯.(*贝特约哈尼斯Bete_yohanisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bete_yohanis",
		TitleName: "贝特约哈尼斯",
		TitleCode: "b_bete_yohanis",
	}
}
