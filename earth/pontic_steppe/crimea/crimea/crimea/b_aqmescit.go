package crimea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克梅斯吉特AqmescitBarony struct {
	feud.BaseBarony
}

var BAqmescit阿克梅斯吉特 feud.Barony = &阿克梅斯吉特AqmescitBarony{}

func init() {
    f := BAqmescit阿克梅斯吉特.(*阿克梅斯吉特AqmescitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aqmescit",
		TitleName: "阿克梅斯吉特",
		TitleCode: "b_aqmescit",
	}
}
