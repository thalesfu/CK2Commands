package kursk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔斯克KurskBarony struct {
	feud.BaseBarony
}

var BKursk库尔斯克 feud.Barony = &库尔斯克KurskBarony{}

func init() {
    f := BKursk库尔斯克.(*库尔斯克KurskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kursk",
		TitleName: "库尔斯克",
		TitleCode: "b_kursk",
	}
}
