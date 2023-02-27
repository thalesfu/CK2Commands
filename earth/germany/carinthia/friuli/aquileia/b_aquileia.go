package aquileia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿奎莱亚AquileiaBarony struct {
	feud.BaseBarony
}

var BAquileia阿奎莱亚 feud.Barony = &阿奎莱亚AquileiaBarony{}

func init() {
    f := BAquileia阿奎莱亚.(*阿奎莱亚AquileiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aquileia",
		TitleName: "阿奎莱亚",
		TitleCode: "b_aquileia",
	}
}
