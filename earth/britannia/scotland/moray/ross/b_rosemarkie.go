package ross

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯马基RosemarkieBarony struct {
	feud.BaseBarony
}

var BRosemarkie罗斯马基 feud.Barony = &罗斯马基RosemarkieBarony{}

func init() {
	f := BRosemarkie罗斯马基.(*罗斯马基RosemarkieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rosemarkie",
		TitleName: "罗斯马基",
		TitleCode: "b_rosemarkie",
	}
}
