package lukomorie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利奇克KalchikBarony struct {
	feud.BaseBarony
}

var BKalchik卡利奇克 feud.Barony = &卡利奇克KalchikBarony{}

func init() {
    f := BKalchik卡利奇克.(*卡利奇克KalchikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalchik",
		TitleName: "卡利奇克",
		TitleCode: "b_kalchik",
	}
}
