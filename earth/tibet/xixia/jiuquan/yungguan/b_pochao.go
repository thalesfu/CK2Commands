package yungguan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坡巢PochaoBarony struct {
	feud.BaseBarony
}

var BPochao坡巢 feud.Barony = &坡巢PochaoBarony{}

func init() {
    f := BPochao坡巢.(*坡巢PochaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pochao",
		TitleName: "坡巢",
		TitleCode: "b_pochao",
	}
}
