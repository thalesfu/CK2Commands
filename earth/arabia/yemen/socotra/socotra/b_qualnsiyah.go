package socotra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古兰西耶QualnsiyahBarony struct {
	feud.BaseBarony
}

var BQualnsiyah古兰西耶 feud.Barony = &古兰西耶QualnsiyahBarony{}

func init() {
    f := BQualnsiyah古兰西耶.(*古兰西耶QualnsiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qualnsiyah",
		TitleName: "古兰西耶",
		TitleCode: "b_qualnsiyah",
	}
}
