package nellore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦罗诃悉帝KalahastiBarony struct {
	feud.BaseBarony
}

var BKalahasti迦罗诃悉帝 feud.Barony = &迦罗诃悉帝KalahastiBarony{}

func init() {
    f := BKalahasti迦罗诃悉帝.(*迦罗诃悉帝KalahastiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalahasti",
		TitleName: "迦罗诃悉帝",
		TitleCode: "b_kalahasti",
	}
}
