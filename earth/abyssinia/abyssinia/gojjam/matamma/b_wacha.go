package matamma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦豪WachaBarony struct {
	feud.BaseBarony
}

var BWacha瓦豪 feud.Barony = &瓦豪WachaBarony{}

func init() {
    f := BWacha瓦豪.(*瓦豪WachaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wacha",
		TitleName: "瓦豪",
		TitleCode: "b_wacha",
	}
}
