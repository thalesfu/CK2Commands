package khinjali_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俱折罗KujarBarony struct {
	feud.BaseBarony
}

var BKujar俱折罗 feud.Barony = &俱折罗KujarBarony{}

func init() {
    f := BKujar俱折罗.(*俱折罗KujarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kujar",
		TitleName: "俱折罗",
		TitleCode: "b_kujar",
	}
}
