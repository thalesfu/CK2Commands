package lower_volga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜博夫卡DubovkaBarony struct {
	feud.BaseBarony
}

var BDubovka杜博夫卡 feud.Barony = &杜博夫卡DubovkaBarony{}

func init() {
    f := BDubovka杜博夫卡.(*杜博夫卡DubovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dubovka",
		TitleName: "杜博夫卡",
		TitleCode: "b_dubovka",
	}
}
