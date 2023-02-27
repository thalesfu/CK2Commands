package oromieh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰尔德兰ChaldiranBarony struct {
	feud.BaseBarony
}

var BChaldiran恰尔德兰 feud.Barony = &恰尔德兰ChaldiranBarony{}

func init() {
    f := BChaldiran恰尔德兰.(*恰尔德兰ChaldiranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaldiran",
		TitleName: "恰尔德兰",
		TitleCode: "b_chaldiran",
	}
}
