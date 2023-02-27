package treviso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴萨诺BassanoBarony struct {
	feud.BaseBarony
}

var BBassano巴萨诺 feud.Barony = &巴萨诺BassanoBarony{}

func init() {
    f := BBassano巴萨诺.(*巴萨诺BassanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bassano",
		TitleName: "巴萨诺",
		TitleCode: "b_bassano",
	}
}
