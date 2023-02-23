package gwynedd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康威ConwyBarony struct {
	feud.BaseBarony
}

var BConwy康威 feud.Barony = &康威ConwyBarony{}

func init() {
	f := BConwy康威.(*康威ConwyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "conwy",
		TitleName: "康威",
		TitleCode: "b_conwy",
	}
}
