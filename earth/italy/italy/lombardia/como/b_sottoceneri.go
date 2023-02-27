package como

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索托切内里SottoceneriBarony struct {
	feud.BaseBarony
}

var BSottoceneri索托切内里 feud.Barony = &索托切内里SottoceneriBarony{}

func init() {
    f := BSottoceneri索托切内里.(*索托切内里SottoceneriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sottoceneri",
		TitleName: "索托切内里",
		TitleCode: "b_sottoceneri",
	}
}
