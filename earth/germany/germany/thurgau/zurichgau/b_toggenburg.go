package zurichgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吐根堡ToggenburgBarony struct {
	feud.BaseBarony
}

var BToggenburg吐根堡 feud.Barony = &吐根堡ToggenburgBarony{}

func init() {
	f := BToggenburg吐根堡.(*吐根堡ToggenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toggenburg",
		TitleName: "吐根堡",
		TitleCode: "b_toggenburg",
	}
}
