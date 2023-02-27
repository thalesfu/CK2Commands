package lleida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格拉蒙AgramuntBarony struct {
	feud.BaseBarony
}

var BAgramunt阿格拉蒙 feud.Barony = &阿格拉蒙AgramuntBarony{}

func init() {
    f := BAgramunt阿格拉蒙.(*阿格拉蒙AgramuntBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agramunt",
		TitleName: "阿格拉蒙",
		TitleCode: "b_agramunt",
	}
}
