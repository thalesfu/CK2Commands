package quena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢克索LuxorBarony struct {
	feud.BaseBarony
}

var BLuxor卢克索 feud.Barony = &卢克索LuxorBarony{}

func init() {
    f := BLuxor卢克索.(*卢克索LuxorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luxor",
		TitleName: "卢克索",
		TitleCode: "b_luxor",
	}
}
