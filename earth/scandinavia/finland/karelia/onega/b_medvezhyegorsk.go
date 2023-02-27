package onega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔胡梅基MedvezhyegorskBarony struct {
	feud.BaseBarony
}

var BMedvezhyegorsk卡尔胡梅基 feud.Barony = &卡尔胡梅基MedvezhyegorskBarony{}

func init() {
    f := BMedvezhyegorsk卡尔胡梅基.(*卡尔胡梅基MedvezhyegorskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medvezhyegorsk",
		TitleName: "卡尔胡梅基",
		TitleCode: "b_medvezhyegorsk",
	}
}
