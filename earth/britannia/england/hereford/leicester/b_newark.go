package leicester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纽瓦克NewarkBarony struct {
	feud.BaseBarony
}

var BNewark纽瓦克 feud.Barony = &纽瓦克NewarkBarony{}

func init() {
	f := BNewark纽瓦克.(*纽瓦克NewarkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "newark",
		TitleName: "纽瓦克",
		TitleCode: "b_newark",
	}
}
