package calatayud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉马德亚拉贡AlhamadearagonBarony struct {
	feud.BaseBarony
}

var BAlhamadearagon阿拉马德亚拉贡 feud.Barony = &阿拉马德亚拉贡AlhamadearagonBarony{}

func init() {
    f := BAlhamadearagon阿拉马德亚拉贡.(*阿拉马德亚拉贡AlhamadearagonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alhamadearagon",
		TitleName: "阿拉马德亚拉贡",
		TitleCode: "b_alhamadearagon",
	}
}
