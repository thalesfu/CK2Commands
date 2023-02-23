package leicester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺丁汉NottinghamBarony struct {
	feud.BaseBarony
}

var BNottingham诺丁汉 feud.Barony = &诺丁汉NottinghamBarony{}

func init() {
	f := BNottingham诺丁汉.(*诺丁汉NottinghamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nottingham",
		TitleName: "诺丁汉",
		TitleCode: "b_nottingham",
	}
}
