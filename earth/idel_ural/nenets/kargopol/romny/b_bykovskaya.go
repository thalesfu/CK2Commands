package romny

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝科夫斯卡亚BykovskayaBarony struct {
	feud.BaseBarony
}

var BBykovskaya贝科夫斯卡亚 feud.Barony = &贝科夫斯卡亚BykovskayaBarony{}

func init() {
    f := BBykovskaya贝科夫斯卡亚.(*贝科夫斯卡亚BykovskayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bykovskaya",
		TitleName: "贝科夫斯卡亚",
		TitleCode: "b_bykovskaya",
	}
}
