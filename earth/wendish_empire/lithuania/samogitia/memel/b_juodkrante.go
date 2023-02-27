package memel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤奥德克兰特JuodkranteBarony struct {
	feud.BaseBarony
}

var BJuodkrante尤奥德克兰特 feud.Barony = &尤奥德克兰特JuodkranteBarony{}

func init() {
    f := BJuodkrante尤奥德克兰特.(*尤奥德克兰特JuodkranteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "juodkrante",
		TitleName: "尤奥德克兰特",
		TitleCode: "b_juodkrante",
	}
}
