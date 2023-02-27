package albania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡巴莱KabalaBarony struct {
	feud.BaseBarony
}

var BKabala卡巴莱 feud.Barony = &卡巴莱KabalaBarony{}

func init() {
    f := BKabala卡巴莱.(*卡巴莱KabalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kabala",
		TitleName: "卡巴莱",
		TitleCode: "b_kabala",
	}
}
