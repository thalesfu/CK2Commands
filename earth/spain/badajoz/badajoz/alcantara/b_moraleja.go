package alcantara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫拉莱哈MoralejaBarony struct {
	feud.BaseBarony
}

var BMoraleja莫拉莱哈 feud.Barony = &莫拉莱哈MoralejaBarony{}

func init() {
	f := BMoraleja莫拉莱哈.(*莫拉莱哈MoralejaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moraleja",
		TitleName: "莫拉莱哈",
		TitleCode: "b_moraleja",
	}
}
