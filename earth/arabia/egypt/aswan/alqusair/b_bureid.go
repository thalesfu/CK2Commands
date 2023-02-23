package alqusair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布赖德BureidBarony struct {
	feud.BaseBarony
}

var BBureid布赖德 feud.Barony = &布赖德BureidBarony{}

func init() {
	f := BBureid布赖德.(*布赖德BureidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bureid",
		TitleName: "布赖德",
		TitleCode: "b_bureid",
	}
}
