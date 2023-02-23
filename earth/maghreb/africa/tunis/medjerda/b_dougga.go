package medjerda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙格DouggaBarony struct {
	feud.BaseBarony
}

var BDougga沙格 feud.Barony = &沙格DouggaBarony{}

func init() {
	f := BDougga沙格.(*沙格DouggaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dougga",
		TitleName: "沙格",
		TitleCode: "b_dougga",
	}
}
