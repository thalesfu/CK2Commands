package narke

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约克斯霍尔姆GoksholmBarony struct {
	feud.BaseBarony
}

var BGoksholm约克斯霍尔姆 feud.Barony = &约克斯霍尔姆GoksholmBarony{}

func init() {
	f := BGoksholm约克斯霍尔姆.(*约克斯霍尔姆GoksholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goksholm",
		TitleName: "约克斯霍尔姆",
		TitleCode: "b_goksholm",
	}
}
