package more

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图绍斯TorsasBarony struct {
	feud.BaseBarony
}

var BTorsas图绍斯 feud.Barony = &图绍斯TorsasBarony{}

func init() {
    f := BTorsas图绍斯.(*图绍斯TorsasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torsas",
		TitleName: "图绍斯",
		TitleCode: "b_torsas",
	}
}
