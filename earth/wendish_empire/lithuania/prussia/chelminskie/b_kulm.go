package chelminskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃扎KulmBarony struct {
	feud.BaseBarony
}

var BKulm沃扎 feud.Barony = &沃扎KulmBarony{}

func init() {
    f := BKulm沃扎.(*沃扎KulmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kulm",
		TitleName: "沃扎",
		TitleCode: "b_kulm",
	}
}
