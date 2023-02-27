package azerbaijan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴斯塔姆BastamBarony struct {
	feud.BaseBarony
}

var BBastam巴斯塔姆 feud.Barony = &巴斯塔姆BastamBarony{}

func init() {
    f := BBastam巴斯塔姆.(*巴斯塔姆BastamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bastam",
		TitleName: "巴斯塔姆",
		TitleCode: "b_bastam",
	}
}
