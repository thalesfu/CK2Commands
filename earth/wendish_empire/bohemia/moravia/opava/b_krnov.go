package opava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔诺夫KrnovBarony struct {
	feud.BaseBarony
}

var BKrnov克尔诺夫 feud.Barony = &克尔诺夫KrnovBarony{}

func init() {
    f := BKrnov克尔诺夫.(*克尔诺夫KrnovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krnov",
		TitleName: "克尔诺夫",
		TitleCode: "b_krnov",
	}
}
