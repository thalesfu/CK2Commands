package terkhiin_tsagaan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔欣查干Terkhiin_tsagaanBarony struct {
	feud.BaseBarony
}

var BTerkhiin_tsagaan特尔欣查干 feud.Barony = &特尔欣查干Terkhiin_tsagaanBarony{}

func init() {
    f := BTerkhiin_tsagaan特尔欣查干.(*特尔欣查干Terkhiin_tsagaanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terkhiin_tsagaan",
		TitleName: "特尔欣查干",
		TitleCode: "b_terkhiin_tsagaan",
	}
}
