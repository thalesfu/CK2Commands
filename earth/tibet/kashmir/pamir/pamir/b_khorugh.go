package pamir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍罗格KhorughBarony struct {
	feud.BaseBarony
}

var BKhorugh霍罗格 feud.Barony = &霍罗格KhorughBarony{}

func init() {
    f := BKhorugh霍罗格.(*霍罗格KhorughBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorugh",
		TitleName: "霍罗格",
		TitleCode: "b_khorugh",
	}
}
