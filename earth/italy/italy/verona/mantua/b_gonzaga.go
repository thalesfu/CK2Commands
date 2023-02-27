package mantua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡扎加GonzagaBarony struct {
	feud.BaseBarony
}

var BGonzaga贡扎加 feud.Barony = &贡扎加GonzagaBarony{}

func init() {
    f := BGonzaga贡扎加.(*贡扎加GonzagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gonzaga",
		TitleName: "贡扎加",
		TitleCode: "b_gonzaga",
	}
}
