package nizhny_novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克斯托沃KstovoBarony struct {
	feud.BaseBarony
}

var BKstovo克斯托沃 feud.Barony = &克斯托沃KstovoBarony{}

func init() {
    f := BKstovo克斯托沃.(*克斯托沃KstovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kstovo",
		TitleName: "克斯托沃",
		TitleCode: "b_kstovo",
	}
}
