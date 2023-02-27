package sarkel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢米卡拉科尔斯克SemikarakorskBarony struct {
	feud.BaseBarony
}

var BSemikarakorsk谢米卡拉科尔斯克 feud.Barony = &谢米卡拉科尔斯克SemikarakorskBarony{}

func init() {
    f := BSemikarakorsk谢米卡拉科尔斯克.(*谢米卡拉科尔斯克SemikarakorskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "semikarakorsk",
		TitleName: "谢米卡拉科尔斯克",
		TitleCode: "b_semikarakorsk",
	}
}
