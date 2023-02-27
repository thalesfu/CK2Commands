package penugonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿难多补罗AnantapurBarony struct {
	feud.BaseBarony
}

var BAnantapur阿难多补罗 feud.Barony = &阿难多补罗AnantapurBarony{}

func init() {
    f := BAnantapur阿难多补罗.(*阿难多补罗AnantapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anantapur",
		TitleName: "阿难多补罗",
		TitleCode: "b_anantapur",
	}
}
