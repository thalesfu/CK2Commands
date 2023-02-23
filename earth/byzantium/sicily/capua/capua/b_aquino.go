package capua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿奎诺AquinoBarony struct {
	feud.BaseBarony
}

var BAquino阿奎诺 feud.Barony = &阿奎诺AquinoBarony{}

func init() {
	f := BAquino阿奎诺.(*阿奎诺AquinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aquino",
		TitleName: "阿奎诺",
		TitleCode: "b_aquino",
	}
}
