package bhera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉迦卡GirjakhBarony struct {
	feud.BaseBarony
}

var BGirjakh吉迦卡 feud.Barony = &吉迦卡GirjakhBarony{}

func init() {
	f := BGirjakh吉迦卡.(*吉迦卡GirjakhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "girjakh",
		TitleName: "吉迦卡",
		TitleCode: "b_girjakh",
	}
}
