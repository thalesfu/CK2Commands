package ajayameru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿阇耶迷卢AjayameruBarony struct {
	feud.BaseBarony
}

var BAjayameru阿阇耶迷卢 feud.Barony = &阿阇耶迷卢AjayameruBarony{}

func init() {
    f := BAjayameru阿阇耶迷卢.(*阿阇耶迷卢AjayameruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ajayameru",
		TitleName: "阿阇耶迷卢",
		TitleCode: "b_ajayameru",
	}
}
