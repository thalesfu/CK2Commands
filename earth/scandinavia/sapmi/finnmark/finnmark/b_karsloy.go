package finnmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔绥KarsloyBarony struct {
	feud.BaseBarony
}

var BKarsloy卡尔绥 feud.Barony = &卡尔绥KarsloyBarony{}

func init() {
    f := BKarsloy卡尔绥.(*卡尔绥KarsloyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karsloy",
		TitleName: "卡尔绥",
		TitleCode: "b_karsloy",
	}
}
