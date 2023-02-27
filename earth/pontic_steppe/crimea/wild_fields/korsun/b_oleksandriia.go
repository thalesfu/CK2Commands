package korsun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚历山德里亚OleksandriiaBarony struct {
	feud.BaseBarony
}

var BOleksandriia亚历山德里亚 feud.Barony = &亚历山德里亚OleksandriiaBarony{}

func init() {
    f := BOleksandriia亚历山德里亚.(*亚历山德里亚OleksandriiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oleksandriia",
		TitleName: "亚历山德里亚",
		TitleCode: "b_oleksandriia",
	}
}
