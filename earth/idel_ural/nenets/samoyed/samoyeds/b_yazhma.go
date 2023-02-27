package samoyeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚日马YazhmaBarony struct {
	feud.BaseBarony
}

var BYazhma亚日马 feud.Barony = &亚日马YazhmaBarony{}

func init() {
    f := BYazhma亚日马.(*亚日马YazhmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yazhma",
		TitleName: "亚日马",
		TitleCode: "b_yazhma",
	}
}
