package tregor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷吉耶TreguierBarony struct {
	feud.BaseBarony
}

var BTreguier特雷吉耶 feud.Barony = &特雷吉耶TreguierBarony{}

func init() {
	f := BTreguier特雷吉耶.(*特雷吉耶TreguierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "treguier",
		TitleName: "特雷吉耶",
		TitleCode: "b_treguier",
	}
}
