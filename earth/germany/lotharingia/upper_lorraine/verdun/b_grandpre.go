package verdun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格朗普雷GrandpreBarony struct {
	feud.BaseBarony
}

var BGrandpre格朗普雷 feud.Barony = &格朗普雷GrandpreBarony{}

func init() {
    f := BGrandpre格朗普雷.(*格朗普雷GrandpreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grandpre",
		TitleName: "格朗普雷",
		TitleCode: "b_grandpre",
	}
}
