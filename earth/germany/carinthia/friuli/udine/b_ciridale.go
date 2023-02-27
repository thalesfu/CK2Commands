package udine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇维达莱CiridaleBarony struct {
	feud.BaseBarony
}

var BCiridale奇维达莱 feud.Barony = &奇维达莱CiridaleBarony{}

func init() {
    f := BCiridale奇维达莱.(*奇维达莱CiridaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ciridale",
		TitleName: "奇维达莱",
		TitleCode: "b_ciridale",
	}
}
