package halaban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷杜木ArradumBarony struct {
	feud.BaseBarony
}

var BArradum雷杜木 feud.Barony = &雷杜木ArradumBarony{}

func init() {
    f := BArradum雷杜木.(*雷杜木ArradumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arradum",
		TitleName: "雷杜木",
		TitleCode: "b_arradum",
	}
}
