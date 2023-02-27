package tobol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔加特卡SargatkaBarony struct {
	feud.BaseBarony
}

var BSargatka萨尔加特卡 feud.Barony = &萨尔加特卡SargatkaBarony{}

func init() {
    f := BSargatka萨尔加特卡.(*萨尔加特卡SargatkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sargatka",
		TitleName: "萨尔加特卡",
		TitleCode: "b_sargatka",
	}
}
