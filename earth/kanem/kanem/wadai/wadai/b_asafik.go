package wadai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿萨菲克AsafikBarony struct {
	feud.BaseBarony
}

var BAsafik阿萨菲克 feud.Barony = &阿萨菲克AsafikBarony{}

func init() {
    f := BAsafik阿萨菲克.(*阿萨菲克AsafikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asafik",
		TitleName: "阿萨菲克",
		TitleCode: "b_asafik",
	}
}
