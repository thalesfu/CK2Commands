package mordva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚瓦斯YavasBarony struct {
	feud.BaseBarony
}

var BYavas亚瓦斯 feud.Barony = &亚瓦斯YavasBarony{}

func init() {
    f := BYavas亚瓦斯.(*亚瓦斯YavasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yavas",
		TitleName: "亚瓦斯",
		TitleCode: "b_yavas",
	}
}
