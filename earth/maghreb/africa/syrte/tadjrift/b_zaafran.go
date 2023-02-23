package tadjrift

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎夫兰ZaafranBarony struct {
	feud.BaseBarony
}

var BZaafran扎夫兰 feud.Barony = &扎夫兰ZaafranBarony{}

func init() {
	f := BZaafran扎夫兰.(*扎夫兰ZaafranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaafran",
		TitleName: "扎夫兰",
		TitleCode: "b_zaafran",
	}
}
