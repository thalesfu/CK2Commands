package vardhamana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿摩伐梨AmarvalliBarony struct {
	feud.BaseBarony
}

var BAmarvalli阿摩伐梨 feud.Barony = &阿摩伐梨AmarvalliBarony{}

func init() {
    f := BAmarvalli阿摩伐梨.(*阿摩伐梨AmarvalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amarvalli",
		TitleName: "阿摩伐梨",
		TitleCode: "b_amarvalli",
	}
}
