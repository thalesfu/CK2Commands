package sikkim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘托克GangdozBarony struct {
	feud.BaseBarony
}

var BGangdoz甘托克 feud.Barony = &甘托克GangdozBarony{}

func init() {
	f := BGangdoz甘托克.(*甘托克GangdozBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gangdoz",
		TitleName: "甘托克",
		TitleCode: "b_gangdoz",
	}
}
