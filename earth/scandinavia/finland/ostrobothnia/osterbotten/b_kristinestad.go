package osterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克里斯蒂南考蓬基KristinestadBarony struct {
	feud.BaseBarony
}

var BKristinestad克里斯蒂南考蓬基 feud.Barony = &克里斯蒂南考蓬基KristinestadBarony{}

func init() {
    f := BKristinestad克里斯蒂南考蓬基.(*克里斯蒂南考蓬基KristinestadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kristinestad",
		TitleName: "克里斯蒂南考蓬基",
		TitleCode: "b_kristinestad",
	}
}
