package aurillac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔拉CarlatBarony struct {
	feud.BaseBarony
}

var BCarlat卡尔拉 feud.Barony = &卡尔拉CarlatBarony{}

func init() {
    f := BCarlat卡尔拉.(*卡尔拉CarlatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carlat",
		TitleName: "卡尔拉",
		TitleCode: "b_carlat",
	}
}
