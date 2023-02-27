package kozelsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科泽利斯克KozelskBarony struct {
	feud.BaseBarony
}

var BKozelsk科泽利斯克 feud.Barony = &科泽利斯克KozelskBarony{}

func init() {
    f := BKozelsk科泽利斯克.(*科泽利斯克KozelskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kozelsk",
		TitleName: "科泽利斯克",
		TitleCode: "b_kozelsk",
	}
}
