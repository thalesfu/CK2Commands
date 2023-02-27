package besancon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬塔利耶PontarlierBarony struct {
	feud.BaseBarony
}

var BPontarlier蓬塔利耶 feud.Barony = &蓬塔利耶PontarlierBarony{}

func init() {
    f := BPontarlier蓬塔利耶.(*蓬塔利耶PontarlierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pontarlier",
		TitleName: "蓬塔利耶",
		TitleCode: "b_pontarlier",
	}
}
