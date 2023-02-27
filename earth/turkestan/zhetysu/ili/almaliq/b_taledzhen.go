package almaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔勒德镇TaledzhenBarony struct {
	feud.BaseBarony
}

var BTaledzhen塔勒德镇 feud.Barony = &塔勒德镇TaledzhenBarony{}

func init() {
    f := BTaledzhen塔勒德镇.(*塔勒德镇TaledzhenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taledzhen",
		TitleName: "塔勒德镇",
		TitleCode: "b_taledzhen",
	}
}
