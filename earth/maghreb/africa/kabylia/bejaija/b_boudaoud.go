package bejaija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布达乌德BoudaoudBarony struct {
	feud.BaseBarony
}

var BBoudaoud布达乌德 feud.Barony = &布达乌德BoudaoudBarony{}

func init() {
	f := BBoudaoud布达乌德.(*布达乌德BoudaoudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boudaoud",
		TitleName: "布达乌德",
		TitleCode: "b_boudaoud",
	}
}
