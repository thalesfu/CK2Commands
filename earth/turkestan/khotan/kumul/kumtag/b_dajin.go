package kumtag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大金DajinBarony struct {
	feud.BaseBarony
}

var BDajin大金 feud.Barony = &大金DajinBarony{}

func init() {
    f := BDajin大金.(*大金DajinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dajin",
		TitleName: "大金",
		TitleCode: "b_dajin",
	}
}
