package atholl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格兰特利GrandtullyBarony struct {
	feud.BaseBarony
}

var BGrandtully格兰特利 feud.Barony = &格兰特利GrandtullyBarony{}

func init() {
	f := BGrandtully格兰特利.(*格兰特利GrandtullyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grandtully",
		TitleName: "格兰特利",
		TitleCode: "b_grandtully",
	}
}
