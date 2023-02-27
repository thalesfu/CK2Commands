package kent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法弗舍姆FavershamBarony struct {
	feud.BaseBarony
}

var BFaversham法弗舍姆 feud.Barony = &法弗舍姆FavershamBarony{}

func init() {
    f := BFaversham法弗舍姆.(*法弗舍姆FavershamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faversham",
		TitleName: "法弗舍姆",
		TitleCode: "b_faversham",
	}
}
