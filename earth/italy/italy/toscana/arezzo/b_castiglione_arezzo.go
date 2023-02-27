package arezzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯蒂廖内Castiglione_arezzoBarony struct {
	feud.BaseBarony
}

var BCastiglione_arezzo卡斯蒂廖内 feud.Barony = &卡斯蒂廖内Castiglione_arezzoBarony{}

func init() {
    f := BCastiglione_arezzo卡斯蒂廖内.(*卡斯蒂廖内Castiglione_arezzoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castiglione_arezzo",
		TitleName: "卡斯蒂廖内",
		TitleCode: "b_castiglione_arezzo",
	}
}
