package kurumba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古雷尔巴西鲁Gourel_bassirouBarony struct {
	feud.BaseBarony
}

var BGourel_bassirou古雷尔巴西鲁 feud.Barony = &古雷尔巴西鲁Gourel_bassirouBarony{}

func init() {
    f := BGourel_bassirou古雷尔巴西鲁.(*古雷尔巴西鲁Gourel_bassirouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gourel_bassirou",
		TitleName: "古雷尔巴西鲁",
		TitleCode: "b_gourel_bassirou",
	}
}
