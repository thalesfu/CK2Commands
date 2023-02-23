package sens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特罗MontereauBarony struct {
	feud.BaseBarony
}

var BMontereau蒙特罗 feud.Barony = &蒙特罗MontereauBarony{}

func init() {
	f := BMontereau蒙特罗.(*蒙特罗MontereauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montereau",
		TitleName: "蒙特罗",
		TitleCode: "b_montereau",
	}
}
