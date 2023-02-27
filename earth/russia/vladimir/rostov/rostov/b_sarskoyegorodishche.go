package rostov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔斯科耶戈罗季谢SarskoyegorodishcheBarony struct {
	feud.BaseBarony
}

var BSarskoyegorodishche萨尔斯科耶戈罗季谢 feud.Barony = &萨尔斯科耶戈罗季谢SarskoyegorodishcheBarony{}

func init() {
    f := BSarskoyegorodishche萨尔斯科耶戈罗季谢.(*萨尔斯科耶戈罗季谢SarskoyegorodishcheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarskoyegorodishche",
		TitleName: "萨尔斯科耶戈罗季谢",
		TitleCode: "b_sarskoyegorodishche",
	}
}
