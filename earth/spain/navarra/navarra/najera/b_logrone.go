package najera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛格罗尼奥LogroneBarony struct {
	feud.BaseBarony
}

var BLogrone洛格罗尼奥 feud.Barony = &洛格罗尼奥LogroneBarony{}

func init() {
	f := BLogrone洛格罗尼奥.(*洛格罗尼奥LogroneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "logrone",
		TitleName: "洛格罗尼奥",
		TitleCode: "b_logrone",
	}
}
