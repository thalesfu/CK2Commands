package nobatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卡沙AkashaBarony struct {
	feud.BaseBarony
}

var BAkasha阿卡沙 feud.Barony = &阿卡沙AkashaBarony{}

func init() {
    f := BAkasha阿卡沙.(*阿卡沙AkashaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akasha",
		TitleName: "阿卡沙",
		TitleCode: "b_akasha",
	}
}
