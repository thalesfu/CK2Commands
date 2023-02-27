package tripolitana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马祖扎哈MazuzahBarony struct {
	feud.BaseBarony
}

var BMazuzah马祖扎哈 feud.Barony = &马祖扎哈MazuzahBarony{}

func init() {
    f := BMazuzah马祖扎哈.(*马祖扎哈MazuzahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazuzah",
		TitleName: "马祖扎哈",
		TitleCode: "b_mazuzah",
	}
}
