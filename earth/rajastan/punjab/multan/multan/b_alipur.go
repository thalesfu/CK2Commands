package multan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里补罗AlipurBarony struct {
	feud.BaseBarony
}

var BAlipur阿里补罗 feud.Barony = &阿里补罗AlipurBarony{}

func init() {
    f := BAlipur阿里补罗.(*阿里补罗AlipurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alipur",
		TitleName: "阿里补罗",
		TitleCode: "b_alipur",
	}
}
