package artois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰鲁阿讷TerwaanBarony struct {
	feud.BaseBarony
}

var BTerwaan泰鲁阿讷 feud.Barony = &泰鲁阿讷TerwaanBarony{}

func init() {
	f := BTerwaan泰鲁阿讷.(*泰鲁阿讷TerwaanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terwaan",
		TitleName: "泰鲁阿讷",
		TitleCode: "b_terwaan",
	}
}
