package gwent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伯加文尼AbergavennyBarony struct {
	feud.BaseBarony
}

var BAbergavenny阿伯加文尼 feud.Barony = &阿伯加文尼AbergavennyBarony{}

func init() {
	f := BAbergavenny阿伯加文尼.(*阿伯加文尼AbergavennyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abergavenny",
		TitleName: "阿伯加文尼",
		TitleCode: "b_abergavenny",
	}
}
