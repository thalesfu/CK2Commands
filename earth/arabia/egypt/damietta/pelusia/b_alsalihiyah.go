package pelusia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨利希耶AlsalihiyahBarony struct {
	feud.BaseBarony
}

var BAlsalihiyah萨利希耶 feud.Barony = &萨利希耶AlsalihiyahBarony{}

func init() {
    f := BAlsalihiyah萨利希耶.(*萨利希耶AlsalihiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alsalihiyah",
		TitleName: "萨利希耶",
		TitleCode: "b_alsalihiyah",
	}
}
