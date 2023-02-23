package zakynthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿利卡纳斯AlikanasBarony struct {
	feud.BaseBarony
}

var BAlikanas阿利卡纳斯 feud.Barony = &阿利卡纳斯AlikanasBarony{}

func init() {
	f := BAlikanas阿利卡纳斯.(*阿利卡纳斯AlikanasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alikanas",
		TitleName: "阿利卡纳斯",
		TitleCode: "b_alikanas",
	}
}
