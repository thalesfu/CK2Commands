package sudovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利斯基亚瓦LiskiavaBarony struct {
	feud.BaseBarony
}

var BLiskiava利斯基亚瓦 feud.Barony = &利斯基亚瓦LiskiavaBarony{}

func init() {
    f := BLiskiava利斯基亚瓦.(*利斯基亚瓦LiskiavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liskiava",
		TitleName: "利斯基亚瓦",
		TitleCode: "b_liskiava",
	}
}
