package ziz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝沙尔BecharBarony struct {
	feud.BaseBarony
}

var BBechar贝沙尔 feud.Barony = &贝沙尔BecharBarony{}

func init() {
	f := BBechar贝沙尔.(*贝沙尔BecharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bechar",
		TitleName: "贝沙尔",
		TitleCode: "b_bechar",
	}
}
