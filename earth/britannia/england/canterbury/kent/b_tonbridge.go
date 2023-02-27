package kent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汤布里奇TonbridgeBarony struct {
	feud.BaseBarony
}

var BTonbridge汤布里奇 feud.Barony = &汤布里奇TonbridgeBarony{}

func init() {
    f := BTonbridge汤布里奇.(*汤布里奇TonbridgeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tonbridge",
		TitleName: "汤布里奇",
		TitleCode: "b_tonbridge",
	}
}
