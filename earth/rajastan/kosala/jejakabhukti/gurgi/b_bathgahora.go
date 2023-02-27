package gurgi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆多迦呼罗BathgahoraBarony struct {
	feud.BaseBarony
}

var BBathgahora婆多迦呼罗 feud.Barony = &婆多迦呼罗BathgahoraBarony{}

func init() {
    f := BBathgahora婆多迦呼罗.(*婆多迦呼罗BathgahoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bathgahora",
		TitleName: "婆多迦呼罗",
		TitleCode: "b_bathgahora",
	}
}
