package damot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔特勒TelteleBarony struct {
	feud.BaseBarony
}

var BTeltele特尔特勒 feud.Barony = &特尔特勒TelteleBarony{}

func init() {
    f := BTeltele特尔特勒.(*特尔特勒TelteleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teltele",
		TitleName: "特尔特勒",
		TitleCode: "b_teltele",
	}
}
