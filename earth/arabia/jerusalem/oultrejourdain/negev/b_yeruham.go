package negev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶鲁汉姆YeruhamBarony struct {
	feud.BaseBarony
}

var BYeruham耶鲁汉姆 feud.Barony = &耶鲁汉姆YeruhamBarony{}

func init() {
    f := BYeruham耶鲁汉姆.(*耶鲁汉姆YeruhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yeruham",
		TitleName: "耶鲁汉姆",
		TitleCode: "b_yeruham",
	}
}
