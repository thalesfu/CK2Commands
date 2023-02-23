package dal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿玛德AmordhBarony struct {
	feud.BaseBarony
}

var BAmordh阿玛德 feud.Barony = &阿玛德AmordhBarony{}

func init() {
	f := BAmordh阿玛德.(*阿玛德AmordhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amordh",
		TitleName: "阿玛德",
		TitleCode: "b_amordh",
	}
}
