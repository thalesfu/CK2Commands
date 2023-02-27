package damot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法沙FashaBarony struct {
	feud.BaseBarony
}

var BFasha法沙 feud.Barony = &法沙FashaBarony{}

func init() {
    f := BFasha法沙.(*法沙FashaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fasha",
		TitleName: "法沙",
		TitleCode: "b_fasha",
	}
}
