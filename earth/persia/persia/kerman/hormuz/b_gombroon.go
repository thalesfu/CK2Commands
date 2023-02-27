package hormuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡布伦GombroonBarony struct {
	feud.BaseBarony
}

var BGombroon贡布伦 feud.Barony = &贡布伦GombroonBarony{}

func init() {
    f := BGombroon贡布伦.(*贡布伦GombroonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gombroon",
		TitleName: "贡布伦",
		TitleCode: "b_gombroon",
	}
}
