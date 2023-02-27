package ochrid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯韦蒂格拉德SvetigradBarony struct {
	feud.BaseBarony
}

var BSvetigrad斯韦蒂格拉德 feud.Barony = &斯韦蒂格拉德SvetigradBarony{}

func init() {
    f := BSvetigrad斯韦蒂格拉德.(*斯韦蒂格拉德SvetigradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svetigrad",
		TitleName: "斯韦蒂格拉德",
		TitleCode: "b_svetigrad",
	}
}
