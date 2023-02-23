package gezira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙瓦夫ShawwafBarony struct {
	feud.BaseBarony
}

var BShawwaf沙瓦夫 feud.Barony = &沙瓦夫ShawwafBarony{}

func init() {
	f := BShawwaf沙瓦夫.(*沙瓦夫ShawwafBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shawwaf",
		TitleName: "沙瓦夫",
		TitleCode: "b_shawwaf",
	}
}
