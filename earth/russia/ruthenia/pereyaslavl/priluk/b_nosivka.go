package priluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺维斯卡NosivkaBarony struct {
	feud.BaseBarony
}

var BNosivka诺维斯卡 feud.Barony = &诺维斯卡NosivkaBarony{}

func init() {
	f := BNosivka诺维斯卡.(*诺维斯卡NosivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nosivka",
		TitleName: "诺维斯卡",
		TitleCode: "b_nosivka",
	}
}
