package nassau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法尔肯施泰因FalkensteinBarony struct {
	feud.BaseBarony
}

var BFalkenstein法尔肯施泰因 feud.Barony = &法尔肯施泰因FalkensteinBarony{}

func init() {
	f := BFalkenstein法尔肯施泰因.(*法尔肯施泰因FalkensteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "falkenstein",
		TitleName: "法尔肯施泰因",
		TitleCode: "b_falkenstein",
	}
}
