package finnveden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮克斯堡PiksborgBarony struct {
	feud.BaseBarony
}

var BPiksborg皮克斯堡 feud.Barony = &皮克斯堡PiksborgBarony{}

func init() {
    f := BPiksborg皮克斯堡.(*皮克斯堡PiksborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piksborg",
		TitleName: "皮克斯堡",
		TitleCode: "b_piksborg",
	}
}
