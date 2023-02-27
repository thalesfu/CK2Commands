package pinsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别廖扎BiarozaBarony struct {
	feud.BaseBarony
}

var BBiaroza别廖扎 feud.Barony = &别廖扎BiarozaBarony{}

func init() {
    f := BBiaroza别廖扎.(*别廖扎BiarozaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biaroza",
		TitleName: "别廖扎",
		TitleCode: "b_biaroza",
	}
}
