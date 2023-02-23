package sasaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼利罗城HorilnagarBarony struct {
	feud.BaseBarony
}

var BHorilnagar呼利罗城 feud.Barony = &呼利罗城HorilnagarBarony{}

func init() {
	f := BHorilnagar呼利罗城.(*呼利罗城HorilnagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horilnagar",
		TitleName: "呼利罗城",
		TitleCode: "b_horilnagar",
	}
}
