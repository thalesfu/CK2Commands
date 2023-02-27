package kumul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洋海YanghaiBarony struct {
	feud.BaseBarony
}

var BYanghai洋海 feud.Barony = &洋海YanghaiBarony{}

func init() {
    f := BYanghai洋海.(*洋海YanghaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yanghai",
		TitleName: "洋海",
		TitleCode: "b_yanghai",
	}
}
