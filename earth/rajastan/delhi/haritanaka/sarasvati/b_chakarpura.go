package sarasvati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 遮迦罗补罗ChakarpuraBarony struct {
	feud.BaseBarony
}

var BChakarpura遮迦罗补罗 feud.Barony = &遮迦罗补罗ChakarpuraBarony{}

func init() {
    f := BChakarpura遮迦罗补罗.(*遮迦罗补罗ChakarpuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chakarpura",
		TitleName: "遮迦罗补罗",
		TitleCode: "b_chakarpura",
	}
}
