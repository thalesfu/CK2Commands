package mosul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔贝拉ArbilBarony struct {
	feud.BaseBarony
}

var BArbil阿尔贝拉 feud.Barony = &阿尔贝拉ArbilBarony{}

func init() {
    f := BArbil阿尔贝拉.(*阿尔贝拉ArbilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arbil",
		TitleName: "阿尔贝拉",
		TitleCode: "b_arbil",
	}
}
