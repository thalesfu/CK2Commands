package ani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣萨尔基斯SurpasdvadzadzinBarony struct {
	feud.BaseBarony
}

var BSurpasdvadzadzin圣萨尔基斯 feud.Barony = &圣萨尔基斯SurpasdvadzadzinBarony{}

func init() {
	f := BSurpasdvadzadzin圣萨尔基斯.(*圣萨尔基斯SurpasdvadzadzinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "surpasdvadzadzin",
		TitleName: "圣萨尔基斯",
		TitleCode: "b_surpasdvadzadzin",
	}
}
