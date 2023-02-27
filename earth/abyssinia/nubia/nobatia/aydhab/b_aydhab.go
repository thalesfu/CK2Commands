package aydhab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伊迪哈卜AydhabBarony struct {
	feud.BaseBarony
}

var BAydhab阿伊迪哈卜 feud.Barony = &阿伊迪哈卜AydhabBarony{}

func init() {
    f := BAydhab阿伊迪哈卜.(*阿伊迪哈卜AydhabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aydhab",
		TitleName: "阿伊迪哈卜",
		TitleCode: "b_aydhab",
	}
}
