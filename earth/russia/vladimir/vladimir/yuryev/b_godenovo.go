package yuryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈杰诺沃GodenovoBarony struct {
	feud.BaseBarony
}

var BGodenovo戈杰诺沃 feud.Barony = &戈杰诺沃GodenovoBarony{}

func init() {
	f := BGodenovo戈杰诺沃.(*戈杰诺沃GodenovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "godenovo",
		TitleName: "戈杰诺沃",
		TitleCode: "b_godenovo",
	}
}
