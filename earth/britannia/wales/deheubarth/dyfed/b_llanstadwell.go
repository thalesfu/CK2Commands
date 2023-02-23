package dyfed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰斯塔德韦尔LlanstadwellBarony struct {
	feud.BaseBarony
}

var BLlanstadwell兰斯塔德韦尔 feud.Barony = &兰斯塔德韦尔LlanstadwellBarony{}

func init() {
	f := BLlanstadwell兰斯塔德韦尔.(*兰斯塔德韦尔LlanstadwellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llanstadwell",
		TitleName: "兰斯塔德韦尔",
		TitleCode: "b_llanstadwell",
	}
}
