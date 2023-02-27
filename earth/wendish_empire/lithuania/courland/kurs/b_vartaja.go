package kurs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔塔亚VartajaBarony struct {
	feud.BaseBarony
}

var BVartaja瓦尔塔亚 feud.Barony = &瓦尔塔亚VartajaBarony{}

func init() {
    f := BVartaja瓦尔塔亚.(*瓦尔塔亚VartajaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vartaja",
		TitleName: "瓦尔塔亚",
		TitleCode: "b_vartaja",
	}
}
