package urbino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福松布罗内FossombroneBarony struct {
	feud.BaseBarony
}

var BFossombrone福松布罗内 feud.Barony = &福松布罗内FossombroneBarony{}

func init() {
	f := BFossombrone福松布罗内.(*福松布罗内FossombroneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fossombrone",
		TitleName: "福松布罗内",
		TitleCode: "b_fossombrone",
	}
}
