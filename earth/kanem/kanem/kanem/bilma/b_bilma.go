package bilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔马BilmaBarony struct {
	feud.BaseBarony
}

var BBilma比尔马 feud.Barony = &比尔马BilmaBarony{}

func init() {
	f := BBilma比尔马.(*比尔马BilmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilma",
		TitleName: "比尔马",
		TitleCode: "b_bilma",
	}
}
