package aprutium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿韦扎诺AvezzanoBarony struct {
	feud.BaseBarony
}

var BAvezzano阿韦扎诺 feud.Barony = &阿韦扎诺AvezzanoBarony{}

func init() {
	f := BAvezzano阿韦扎诺.(*阿韦扎诺AvezzanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avezzano",
		TitleName: "阿韦扎诺",
		TitleCode: "b_avezzano",
	}
}
