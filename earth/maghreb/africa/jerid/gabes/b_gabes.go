package gabes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加贝斯GabesBarony struct {
	feud.BaseBarony
}

var BGabes加贝斯 feud.Barony = &加贝斯GabesBarony{}

func init() {
	f := BGabes加贝斯.(*加贝斯GabesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gabes",
		TitleName: "加贝斯",
		TitleCode: "b_gabes",
	}
}
