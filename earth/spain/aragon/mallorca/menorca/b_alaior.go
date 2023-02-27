package menorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿莱奥尔AlaiorBarony struct {
	feud.BaseBarony
}

var BAlaior阿莱奥尔 feud.Barony = &阿莱奥尔AlaiorBarony{}

func init() {
    f := BAlaior阿莱奥尔.(*阿莱奥尔AlaiorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alaior",
		TitleName: "阿莱奥尔",
		TitleCode: "b_alaior",
	}
}
