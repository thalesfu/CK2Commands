package gabes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞夫提米SteftimiBarony struct {
	feud.BaseBarony
}

var BSteftimi塞夫提米 feud.Barony = &塞夫提米SteftimiBarony{}

func init() {
	f := BSteftimi塞夫提米.(*塞夫提米SteftimiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "steftimi",
		TitleName: "塞夫提米",
		TitleCode: "b_steftimi",
	}
}
