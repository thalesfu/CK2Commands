package suenik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈里斯GorisBarony struct {
	feud.BaseBarony
}

var BGoris戈里斯 feud.Barony = &戈里斯GorisBarony{}

func init() {
	f := BGoris戈里斯.(*戈里斯GorisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goris",
		TitleName: "戈里斯",
		TitleCode: "b_goris",
	}
}
