package kromy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳夫利亚NavlyaBarony struct {
	feud.BaseBarony
}

var BNavlya纳夫利亚 feud.Barony = &纳夫利亚NavlyaBarony{}

func init() {
    f := BNavlya纳夫利亚.(*纳夫利亚NavlyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "navlya",
		TitleName: "纳夫利亚",
		TitleCode: "b_navlya",
	}
}
