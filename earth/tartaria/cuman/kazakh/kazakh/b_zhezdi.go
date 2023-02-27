package kazakh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰兹迪ZhezdiBarony struct {
	feud.BaseBarony
}

var BZhezdi杰兹迪 feud.Barony = &杰兹迪ZhezdiBarony{}

func init() {
    f := BZhezdi杰兹迪.(*杰兹迪ZhezdiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhezdi",
		TitleName: "杰兹迪",
		TitleCode: "b_zhezdi",
	}
}
