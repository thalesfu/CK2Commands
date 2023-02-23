package zabid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宰比德ZabidBarony struct {
	feud.BaseBarony
}

var BZabid宰比德 feud.Barony = &宰比德ZabidBarony{}

func init() {
	f := BZabid宰比德.(*宰比德ZabidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zabid",
		TitleName: "宰比德",
		TitleCode: "b_zabid",
	}
}
