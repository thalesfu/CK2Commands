package monemvasia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卡迪亚ArkadiaBarony struct {
	feud.BaseBarony
}

var BArkadia阿卡迪亚 feud.Barony = &阿卡迪亚ArkadiaBarony{}

func init() {
    f := BArkadia阿卡迪亚.(*阿卡迪亚ArkadiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arkadia",
		TitleName: "阿卡迪亚",
		TitleCode: "b_arkadia",
	}
}
