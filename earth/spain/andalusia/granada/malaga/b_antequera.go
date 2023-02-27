package malaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安特克拉AntequeraBarony struct {
	feud.BaseBarony
}

var BAntequera安特克拉 feud.Barony = &安特克拉AntequeraBarony{}

func init() {
    f := BAntequera安特克拉.(*安特克拉AntequeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antequera",
		TitleName: "安特克拉",
		TitleCode: "b_antequera",
	}
}
