package barskhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拔塞干BarskhanBarony struct {
	feud.BaseBarony
}

var BBarskhan拔塞干 feud.Barony = &拔塞干BarskhanBarony{}

func init() {
	f := BBarskhan拔塞干.(*拔塞干BarskhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barskhan",
		TitleName: "拔塞干",
		TitleCode: "b_barskhan",
	}
}
