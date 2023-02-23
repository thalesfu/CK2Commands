package massawa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿杜利斯AdulisBarony struct {
	feud.BaseBarony
}

var BAdulis阿杜利斯 feud.Barony = &阿杜利斯AdulisBarony{}

func init() {
	f := BAdulis阿杜利斯.(*阿杜利斯AdulisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adulis",
		TitleName: "阿杜利斯",
		TitleCode: "b_adulis",
	}
}
