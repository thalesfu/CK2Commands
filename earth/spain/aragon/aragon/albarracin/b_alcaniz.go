package albarracin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔卡尼斯AlcanizBarony struct {
	feud.BaseBarony
}

var BAlcaniz阿尔卡尼斯 feud.Barony = &阿尔卡尼斯AlcanizBarony{}

func init() {
    f := BAlcaniz阿尔卡尼斯.(*阿尔卡尼斯AlcanizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcaniz",
		TitleName: "阿尔卡尼斯",
		TitleCode: "b_alcaniz",
	}
}
