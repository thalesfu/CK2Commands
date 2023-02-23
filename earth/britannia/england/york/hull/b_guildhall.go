package hull

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉尔德霍尔GuildhallBarony struct {
	feud.BaseBarony
}

var BGuildhall吉尔德霍尔 feud.Barony = &吉尔德霍尔GuildhallBarony{}

func init() {
	f := BGuildhall吉尔德霍尔.(*吉尔德霍尔GuildhallBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guildhall",
		TitleName: "吉尔德霍尔",
		TitleCode: "b_guildhall",
	}
}
