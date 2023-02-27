package atlas_mnt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰勒法DjelfaBarony struct {
	feud.BaseBarony
}

var BDjelfa杰勒法 feud.Barony = &杰勒法DjelfaBarony{}

func init() {
    f := BDjelfa杰勒法.(*杰勒法DjelfaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djelfa",
		TitleName: "杰勒法",
		TitleCode: "b_djelfa",
	}
}
