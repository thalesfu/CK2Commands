package kafirkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尼诃特ManiotBarony struct {
	feud.BaseBarony
}

var BManiot马尼诃特 feud.Barony = &马尼诃特ManiotBarony{}

func init() {
	f := BManiot马尼诃特.(*马尼诃特ManiotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maniot",
		TitleName: "马尼诃特",
		TitleCode: "b_maniot",
	}
}
