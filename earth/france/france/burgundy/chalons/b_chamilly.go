package chalons

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙米伊ChamillyBarony struct {
	feud.BaseBarony
}

var BChamilly沙米伊 feud.Barony = &沙米伊ChamillyBarony{}

func init() {
    f := BChamilly沙米伊.(*沙米伊ChamillyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chamilly",
		TitleName: "沙米伊",
		TitleCode: "b_chamilly",
	}
}
