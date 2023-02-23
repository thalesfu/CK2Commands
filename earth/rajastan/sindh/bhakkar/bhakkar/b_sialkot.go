package bhakkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 犀耶罗拘吒SialkotBarony struct {
	feud.BaseBarony
}

var BSialkot犀耶罗拘吒 feud.Barony = &犀耶罗拘吒SialkotBarony{}

func init() {
	f := BSialkot犀耶罗拘吒.(*犀耶罗拘吒SialkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sialkot",
		TitleName: "犀耶罗拘吒",
		TitleCode: "b_sialkot",
	}
}
