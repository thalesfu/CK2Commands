package tsilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁江卡RudyankaBarony struct {
	feud.BaseBarony
}

var BRudyanka鲁江卡 feud.Barony = &鲁江卡RudyankaBarony{}

func init() {
    f := BRudyanka鲁江卡.(*鲁江卡RudyankaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rudyanka",
		TitleName: "鲁江卡",
		TitleCode: "b_rudyanka",
	}
}
