package brescia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯蒂廖内CastiglioneBarony struct {
	feud.BaseBarony
}

var BCastiglione卡斯蒂廖内 feud.Barony = &卡斯蒂廖内CastiglioneBarony{}

func init() {
    f := BCastiglione卡斯蒂廖内.(*卡斯蒂廖内CastiglioneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castiglione",
		TitleName: "卡斯蒂廖内",
		TitleCode: "b_castiglione",
	}
}
