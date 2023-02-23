package gondar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆巴瑟AmbasselBarony struct {
	feud.BaseBarony
}

var BAmbassel阿姆巴瑟 feud.Barony = &阿姆巴瑟AmbasselBarony{}

func init() {
	f := BAmbassel阿姆巴瑟.(*阿姆巴瑟AmbasselBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ambassel",
		TitleName: "阿姆巴瑟",
		TitleCode: "b_ambassel",
	}
}
