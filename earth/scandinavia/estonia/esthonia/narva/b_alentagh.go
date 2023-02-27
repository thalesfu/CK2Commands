package narva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿林塔加AlentaghBarony struct {
	feud.BaseBarony
}

var BAlentagh阿林塔加 feud.Barony = &阿林塔加AlentaghBarony{}

func init() {
    f := BAlentagh阿林塔加.(*阿林塔加AlentaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alentagh",
		TitleName: "阿林塔加",
		TitleCode: "b_alentagh",
	}
}
