package pecs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩奇瓦劳德PecsvaradBarony struct {
	feud.BaseBarony
}

var BPecsvarad佩奇瓦劳德 feud.Barony = &佩奇瓦劳德PecsvaradBarony{}

func init() {
    f := BPecsvarad佩奇瓦劳德.(*佩奇瓦劳德PecsvaradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pecsvarad",
		TitleName: "佩奇瓦劳德",
		TitleCode: "b_pecsvarad",
	}
}
