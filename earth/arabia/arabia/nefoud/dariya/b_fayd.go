package dariya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费德FaydBarony struct {
	feud.BaseBarony
}

var BFayd费德 feud.Barony = &费德FaydBarony{}

func init() {
    f := BFayd费德.(*费德FaydBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fayd",
		TitleName: "费德",
		TitleCode: "b_fayd",
	}
}
