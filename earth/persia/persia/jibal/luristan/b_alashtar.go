package luristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉什塔尔AlashtarBarony struct {
	feud.BaseBarony
}

var BAlashtar阿拉什塔尔 feud.Barony = &阿拉什塔尔AlashtarBarony{}

func init() {
    f := BAlashtar阿拉什塔尔.(*阿拉什塔尔AlashtarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alashtar",
		TitleName: "阿拉什塔尔",
		TitleCode: "b_alashtar",
	}
}
