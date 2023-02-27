package gobir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦贾WajaBarony struct {
	feud.BaseBarony
}

var BWaja瓦贾 feud.Barony = &瓦贾WajaBarony{}

func init() {
    f := BWaja瓦贾.(*瓦贾WajaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waja",
		TitleName: "瓦贾",
		TitleCode: "b_waja",
	}
}
