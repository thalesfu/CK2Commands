package braunschweig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔德克WaldeckBarony struct {
	feud.BaseBarony
}

var BWaldeck瓦尔德克 feud.Barony = &瓦尔德克WaldeckBarony{}

func init() {
	f := BWaldeck瓦尔德克.(*瓦尔德克WaldeckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waldeck",
		TitleName: "瓦尔德克",
		TitleCode: "b_waldeck",
	}
}
