package lusignan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔福朗ConfolensBarony struct {
	feud.BaseBarony
}

var BConfolens孔福朗 feud.Barony = &孔福朗ConfolensBarony{}

func init() {
	f := BConfolens孔福朗.(*孔福朗ConfolensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "confolens",
		TitleName: "孔福朗",
		TitleCode: "b_confolens",
	}
}
