package atlas_mnt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因里什AinrichBarony struct {
	feud.BaseBarony
}

var BAinrich艾因里什 feud.Barony = &艾因里什AinrichBarony{}

func init() {
    f := BAinrich艾因里什.(*艾因里什AinrichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ainrich",
		TitleName: "艾因里什",
		TitleCode: "b_ainrich",
	}
}
