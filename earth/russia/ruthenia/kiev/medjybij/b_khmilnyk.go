package medjybij

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫梅利尼克KhmilnykBarony struct {
	feud.BaseBarony
}

var BKhmilnyk赫梅利尼克 feud.Barony = &赫梅利尼克KhmilnykBarony{}

func init() {
	f := BKhmilnyk赫梅利尼克.(*赫梅利尼克KhmilnykBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khmilnyk",
		TitleName: "赫梅利尼克",
		TitleCode: "b_khmilnyk",
	}
}
