package kurmanchal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尸拘吒Askot_kurmanchalBarony struct {
	feud.BaseBarony
}

var BAskot_kurmanchal阿尸拘吒 feud.Barony = &阿尸拘吒Askot_kurmanchalBarony{}

func init() {
    f := BAskot_kurmanchal阿尸拘吒.(*阿尸拘吒Askot_kurmanchalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "askot_kurmanchal",
		TitleName: "阿尸拘吒",
		TitleCode: "b_askot_kurmanchal",
	}
}
