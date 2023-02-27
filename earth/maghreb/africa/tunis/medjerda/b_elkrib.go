package medjerda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯里卜ElkribBarony struct {
	feud.BaseBarony
}

var BElkrib凯里卜 feud.Barony = &凯里卜ElkribBarony{}

func init() {
    f := BElkrib凯里卜.(*凯里卜ElkribBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elkrib",
		TitleName: "凯里卜",
		TitleCode: "b_elkrib",
	}
}
