package satakunta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希腾哈里尤HiittenharjuBarony struct {
	feud.BaseBarony
}

var BHiittenharju希腾哈里尤 feud.Barony = &希腾哈里尤HiittenharjuBarony{}

func init() {
    f := BHiittenharju希腾哈里尤.(*希腾哈里尤HiittenharjuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hiittenharju",
		TitleName: "希腾哈里尤",
		TitleCode: "b_hiittenharju",
	}
}
