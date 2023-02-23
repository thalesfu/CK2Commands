package weimar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔肯里德WalkenriedBarony struct {
	feud.BaseBarony
}

var BWalkenried瓦尔肯里德 feud.Barony = &瓦尔肯里德WalkenriedBarony{}

func init() {
	f := BWalkenried瓦尔肯里德.(*瓦尔肯里德WalkenriedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "walkenried",
		TitleName: "瓦尔肯里德",
		TitleCode: "b_walkenried",
	}
}
