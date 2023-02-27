package khaybar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豪拉HauraBarony struct {
	feud.BaseBarony
}

var BHaura豪拉 feud.Barony = &豪拉HauraBarony{}

func init() {
    f := BHaura豪拉.(*豪拉HauraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haura",
		TitleName: "豪拉",
		TitleCode: "b_haura",
	}
}
