package yelets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫列夫诺耶KhlevnoyeBarony struct {
	feud.BaseBarony
}

var BKhlevnoye赫列夫诺耶 feud.Barony = &赫列夫诺耶KhlevnoyeBarony{}

func init() {
	f := BKhlevnoye赫列夫诺耶.(*赫列夫诺耶KhlevnoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khlevnoye",
		TitleName: "赫列夫诺耶",
		TitleCode: "b_khlevnoye",
	}
}
