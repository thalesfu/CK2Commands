package halland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法尔肯贝里FalkenbergBarony struct {
	feud.BaseBarony
}

var BFalkenberg法尔肯贝里 feud.Barony = &法尔肯贝里FalkenbergBarony{}

func init() {
	f := BFalkenberg法尔肯贝里.(*法尔肯贝里FalkenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "falkenberg",
		TitleName: "法尔肯贝里",
		TitleCode: "b_falkenberg",
	}
}
