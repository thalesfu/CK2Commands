package rahbah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布桑BusanBarony struct {
	feud.BaseBarony
}

var BBusan布桑 feud.Barony = &布桑BusanBarony{}

func init() {
    f := BBusan布桑.(*布桑BusanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "busan",
		TitleName: "布桑",
		TitleCode: "b_busan",
	}
}
