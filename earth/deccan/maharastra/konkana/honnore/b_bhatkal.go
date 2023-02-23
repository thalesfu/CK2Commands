package honnore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆吒迦罗BhatkalBarony struct {
	feud.BaseBarony
}

var BBhatkal婆吒迦罗 feud.Barony = &婆吒迦罗BhatkalBarony{}

func init() {
	f := BBhatkal婆吒迦罗.(*婆吒迦罗BhatkalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhatkal",
		TitleName: "婆吒迦罗",
		TitleCode: "b_bhatkal",
	}
}
