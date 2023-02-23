package pratishthana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波罗提瑟吒那PratishthanaBarony struct {
	feud.BaseBarony
}

var BPratishthana波罗提瑟吒那 feud.Barony = &波罗提瑟吒那PratishthanaBarony{}

func init() {
	f := BPratishthana波罗提瑟吒那.(*波罗提瑟吒那PratishthanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pratishthana",
		TitleName: "波罗提瑟吒那",
		TitleCode: "b_pratishthana",
	}
}
