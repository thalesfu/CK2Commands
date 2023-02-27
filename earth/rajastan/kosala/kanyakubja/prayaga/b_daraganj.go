package prayaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀罗犍阇DaraganjBarony struct {
	feud.BaseBarony
}

var BDaraganj陀罗犍阇 feud.Barony = &陀罗犍阇DaraganjBarony{}

func init() {
    f := BDaraganj陀罗犍阇.(*陀罗犍阇DaraganjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daraganj",
		TitleName: "陀罗犍阇",
		TitleCode: "b_daraganj",
	}
}
