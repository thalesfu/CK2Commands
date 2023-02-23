package njudung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 永阿LjungaBarony struct {
	feud.BaseBarony
}

var BLjunga永阿 feud.Barony = &永阿LjungaBarony{}

func init() {
	f := BLjunga永阿.(*永阿LjungaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ljunga",
		TitleName: "永阿",
		TitleCode: "b_ljunga",
	}
}
