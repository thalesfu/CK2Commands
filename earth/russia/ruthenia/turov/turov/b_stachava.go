package turov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔霍沃StachavaBarony struct {
	feud.BaseBarony
}

var BStachava斯塔霍沃 feud.Barony = &斯塔霍沃StachavaBarony{}

func init() {
	f := BStachava斯塔霍沃.(*斯塔霍沃StachavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stachava",
		TitleName: "斯塔霍沃",
		TitleCode: "b_stachava",
	}
}
