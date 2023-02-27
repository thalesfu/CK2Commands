package granada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安特克拉AntequaraBarony struct {
	feud.BaseBarony
}

var BAntequara安特克拉 feud.Barony = &安特克拉AntequaraBarony{}

func init() {
    f := BAntequara安特克拉.(*安特克拉AntequaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antequara",
		TitleName: "安特克拉",
		TitleCode: "b_antequara",
	}
}
