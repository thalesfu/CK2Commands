package gauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡呼KumhaulBarony struct {
	feud.BaseBarony
}

var BKumhaul贡呼 feud.Barony = &贡呼KumhaulBarony{}

func init() {
    f := BKumhaul贡呼.(*贡呼KumhaulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumhaul",
		TitleName: "贡呼",
		TitleCode: "b_kumhaul",
	}
}
