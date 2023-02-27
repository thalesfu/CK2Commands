package bjarmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥库洛夫斯基OkulovskyBarony struct {
	feud.BaseBarony
}

var BOkulovsky奥库洛夫斯基 feud.Barony = &奥库洛夫斯基OkulovskyBarony{}

func init() {
    f := BOkulovsky奥库洛夫斯基.(*奥库洛夫斯基OkulovskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "okulovsky",
		TitleName: "奥库洛夫斯基",
		TitleCode: "b_okulovsky",
	}
}
