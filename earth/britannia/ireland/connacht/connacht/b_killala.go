package connacht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基拉拉KillalaBarony struct {
	feud.BaseBarony
}

var BKillala基拉拉 feud.Barony = &基拉拉KillalaBarony{}

func init() {
	f := BKillala基拉拉.(*基拉拉KillalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "killala",
		TitleName: "基拉拉",
		TitleCode: "b_killala",
	}
}
