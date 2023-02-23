package mansura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃罗HalaBarony struct {
	feud.BaseBarony
}

var BHala诃罗 feud.Barony = &诃罗HalaBarony{}

func init() {
	f := BHala诃罗.(*诃罗HalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hala",
		TitleName: "诃罗",
		TitleCode: "b_hala",
	}
}
