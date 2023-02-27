package feher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉德AradBarony struct {
	feud.BaseBarony
}

var BArad阿拉德 feud.Barony = &阿拉德AradBarony{}

func init() {
    f := BArad阿拉德.(*阿拉德AradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arad",
		TitleName: "阿拉德",
		TitleCode: "b_arad",
	}
}
