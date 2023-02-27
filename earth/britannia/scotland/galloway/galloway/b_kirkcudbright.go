package galloway

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柯库布里KirkcudbrightBarony struct {
	feud.BaseBarony
}

var BKirkcudbright柯库布里 feud.Barony = &柯库布里KirkcudbrightBarony{}

func init() {
    f := BKirkcudbright柯库布里.(*柯库布里KirkcudbrightBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirkcudbright",
		TitleName: "柯库布里",
		TitleCode: "b_kirkcudbright",
	}
}
