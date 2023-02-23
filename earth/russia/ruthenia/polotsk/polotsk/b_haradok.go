package polotsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈罗多克HaradokBarony struct {
	feud.BaseBarony
}

var BHaradok戈罗多克 feud.Barony = &戈罗多克HaradokBarony{}

func init() {
	f := BHaradok戈罗多克.(*戈罗多克HaradokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haradok",
		TitleName: "戈罗多克",
		TitleCode: "b_haradok",
	}
}
