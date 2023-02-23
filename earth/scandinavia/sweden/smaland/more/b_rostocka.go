package more

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯托卡RostockaBarony struct {
	feud.BaseBarony
}

var BRostocka罗斯托卡 feud.Barony = &罗斯托卡RostockaBarony{}

func init() {
	f := BRostocka罗斯托卡.(*罗斯托卡RostockaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rostocka",
		TitleName: "罗斯托卡",
		TitleCode: "b_rostocka",
	}
}
