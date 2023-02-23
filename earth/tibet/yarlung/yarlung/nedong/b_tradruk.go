package nedong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昌珠TradrukBarony struct {
	feud.BaseBarony
}

var BTradruk昌珠 feud.Barony = &昌珠TradrukBarony{}

func init() {
	f := BTradruk昌珠.(*昌珠TradrukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tradruk",
		TitleName: "昌珠",
		TitleCode: "b_tradruk",
	}
}
