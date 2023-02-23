package tabriz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿哈尔AharBarony struct {
	feud.BaseBarony
}

var BAhar阿哈尔 feud.Barony = &阿哈尔AharBarony{}

func init() {
	f := BAhar阿哈尔.(*阿哈尔AharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahar",
		TitleName: "阿哈尔",
		TitleCode: "b_ahar",
	}
}
