package boulogne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布洛涅BoulogneBarony struct {
	feud.BaseBarony
}

var BBoulogne布洛涅 feud.Barony = &布洛涅BoulogneBarony{}

func init() {
    f := BBoulogne布洛涅.(*布洛涅BoulogneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boulogne",
		TitleName: "布洛涅",
		TitleCode: "b_boulogne",
	}
}
