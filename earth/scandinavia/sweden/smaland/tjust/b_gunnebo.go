package tjust

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡讷布GunneboBarony struct {
	feud.BaseBarony
}

var BGunnebo贡讷布 feud.Barony = &贡讷布GunneboBarony{}

func init() {
    f := BGunnebo贡讷布.(*贡讷布GunneboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gunnebo",
		TitleName: "贡讷布",
		TitleCode: "b_gunnebo",
	}
}
