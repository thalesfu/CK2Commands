package phiti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 密亨达勒MihintaleBarony struct {
	feud.BaseBarony
}

var BMihintale密亨达勒 feud.Barony = &密亨达勒MihintaleBarony{}

func init() {
	f := BMihintale密亨达勒.(*密亨达勒MihintaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mihintale",
		TitleName: "密亨达勒",
		TitleCode: "b_mihintale",
	}
}
