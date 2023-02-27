package tahoua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古鲁米GurumiBarony struct {
	feud.BaseBarony
}

var BGurumi古鲁米 feud.Barony = &古鲁米GurumiBarony{}

func init() {
    f := BGurumi古鲁米.(*古鲁米GurumiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurumi",
		TitleName: "古鲁米",
		TitleCode: "b_gurumi",
	}
}
