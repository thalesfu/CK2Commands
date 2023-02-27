package vastmanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雪平KopingBarony struct {
	feud.BaseBarony
}

var BKoping雪平 feud.Barony = &雪平KopingBarony{}

func init() {
    f := BKoping雪平.(*雪平KopingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koping",
		TitleName: "雪平",
		TitleCode: "b_koping",
	}
}
