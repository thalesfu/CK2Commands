package annaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安纳巴AnnabahBarony struct {
	feud.BaseBarony
}

var BAnnabah安纳巴 feud.Barony = &安纳巴AnnabahBarony{}

func init() {
    f := BAnnabah安纳巴.(*安纳巴AnnabahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "annabah",
		TitleName: "安纳巴",
		TitleCode: "b_annabah",
	}
}
