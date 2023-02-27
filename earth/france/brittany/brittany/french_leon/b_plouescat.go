package french_leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普卢埃斯卡PlouescatBarony struct {
	feud.BaseBarony
}

var BPlouescat普卢埃斯卡 feud.Barony = &普卢埃斯卡PlouescatBarony{}

func init() {
    f := BPlouescat普卢埃斯卡.(*普卢埃斯卡PlouescatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plouescat",
		TitleName: "普卢埃斯卡",
		TitleCode: "b_plouescat",
	}
}
