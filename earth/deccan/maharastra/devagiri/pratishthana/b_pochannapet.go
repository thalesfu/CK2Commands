package pratishthana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩旃那毗提PochannapetBarony struct {
	feud.BaseBarony
}

var BPochannapet菩旃那毗提 feud.Barony = &菩旃那毗提PochannapetBarony{}

func init() {
	f := BPochannapet菩旃那毗提.(*菩旃那毗提PochannapetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pochannapet",
		TitleName: "菩旃那毗提",
		TitleCode: "b_pochannapet",
	}
}
