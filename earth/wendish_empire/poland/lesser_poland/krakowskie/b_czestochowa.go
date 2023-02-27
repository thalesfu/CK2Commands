package krakowskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琴斯托霍瓦CzestochowaBarony struct {
	feud.BaseBarony
}

var BCzestochowa琴斯托霍瓦 feud.Barony = &琴斯托霍瓦CzestochowaBarony{}

func init() {
    f := BCzestochowa琴斯托霍瓦.(*琴斯托霍瓦CzestochowaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "czestochowa",
		TitleName: "琴斯托霍瓦",
		TitleCode: "b_czestochowa",
	}
}
