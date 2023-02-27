package hardsyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔斯特布罗HolstebroBarony struct {
	feud.BaseBarony
}

var BHolstebro霍尔斯特布罗 feud.Barony = &霍尔斯特布罗HolstebroBarony{}

func init() {
    f := BHolstebro霍尔斯特布罗.(*霍尔斯特布罗HolstebroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "holstebro",
		TitleName: "霍尔斯特布罗",
		TitleCode: "b_holstebro",
	}
}
