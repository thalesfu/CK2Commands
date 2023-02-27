package ubagan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科祖拜KozubayBarony struct {
	feud.BaseBarony
}

var BKozubay科祖拜 feud.Barony = &科祖拜KozubayBarony{}

func init() {
    f := BKozubay科祖拜.(*科祖拜KozubayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kozubay",
		TitleName: "科祖拜",
		TitleCode: "b_kozubay",
	}
}
