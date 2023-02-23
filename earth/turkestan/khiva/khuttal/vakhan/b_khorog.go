package vakhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍罗格KhorogBarony struct {
	feud.BaseBarony
}

var BKhorog霍罗格 feud.Barony = &霍罗格KhorogBarony{}

func init() {
	f := BKhorog霍罗格.(*霍罗格KhorogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorog",
		TitleName: "霍罗格",
		TitleCode: "b_khorog",
	}
}
