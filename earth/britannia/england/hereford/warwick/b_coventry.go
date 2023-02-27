package warwick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考文垂CoventryBarony struct {
	feud.BaseBarony
}

var BCoventry考文垂 feud.Barony = &考文垂CoventryBarony{}

func init() {
    f := BCoventry考文垂.(*考文垂CoventryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coventry",
		TitleName: "考文垂",
		TitleCode: "b_coventry",
	}
}
