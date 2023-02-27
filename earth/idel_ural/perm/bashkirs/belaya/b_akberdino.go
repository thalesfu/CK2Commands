package belaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克别尔季诺AkberdinoBarony struct {
	feud.BaseBarony
}

var BAkberdino阿克别尔季诺 feud.Barony = &阿克别尔季诺AkberdinoBarony{}

func init() {
    f := BAkberdino阿克别尔季诺.(*阿克别尔季诺AkberdinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akberdino",
		TitleName: "阿克别尔季诺",
		TitleCode: "b_akberdino",
	}
}
