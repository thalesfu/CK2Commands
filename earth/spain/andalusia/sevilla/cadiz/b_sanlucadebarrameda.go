package cadiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑卢卡尔德巴拉梅达SanlucadebarramedaBarony struct {
	feud.BaseBarony
}

var BSanlucadebarrameda桑卢卡尔德巴拉梅达 feud.Barony = &桑卢卡尔德巴拉梅达SanlucadebarramedaBarony{}

func init() {
    f := BSanlucadebarrameda桑卢卡尔德巴拉梅达.(*桑卢卡尔德巴拉梅达SanlucadebarramedaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanlucadebarrameda",
		TitleName: "桑卢卡尔德巴拉梅达",
		TitleCode: "b_sanlucadebarrameda",
	}
}
