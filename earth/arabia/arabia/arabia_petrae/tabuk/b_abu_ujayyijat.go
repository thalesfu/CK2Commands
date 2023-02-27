package tabuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布欧杰伊贾特Abu_ujayyijatBarony struct {
	feud.BaseBarony
}

var BAbu_ujayyijat阿布欧杰伊贾特 feud.Barony = &阿布欧杰伊贾特Abu_ujayyijatBarony{}

func init() {
    f := BAbu_ujayyijat阿布欧杰伊贾特.(*阿布欧杰伊贾特Abu_ujayyijatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abu_ujayyijat",
		TitleName: "阿布欧杰伊贾特",
		TitleCode: "b_abu_ujayyijat",
	}
}
