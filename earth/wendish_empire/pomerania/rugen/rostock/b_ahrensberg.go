package rostock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伦斯堡AhrensbergBarony struct {
	feud.BaseBarony
}

var BAhrensberg阿伦斯堡 feud.Barony = &阿伦斯堡AhrensbergBarony{}

func init() {
    f := BAhrensberg阿伦斯堡.(*阿伦斯堡AhrensbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahrensberg",
		TitleName: "阿伦斯堡",
		TitleCode: "b_ahrensberg",
	}
}
