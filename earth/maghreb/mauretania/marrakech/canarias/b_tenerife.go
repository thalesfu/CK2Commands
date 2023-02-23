package canarias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特内里费TenerifeBarony struct {
	feud.BaseBarony
}

var BTenerife特内里费 feud.Barony = &特内里费TenerifeBarony{}

func init() {
	f := BTenerife特内里费.(*特内里费TenerifeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tenerife",
		TitleName: "特内里费",
		TitleCode: "b_tenerife",
	}
}
