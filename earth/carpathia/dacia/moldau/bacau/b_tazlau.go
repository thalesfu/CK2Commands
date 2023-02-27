package bacau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔兹勒乌TazlauBarony struct {
	feud.BaseBarony
}

var BTazlau塔兹勒乌 feud.Barony = &塔兹勒乌TazlauBarony{}

func init() {
    f := BTazlau塔兹勒乌.(*塔兹勒乌TazlauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tazlau",
		TitleName: "塔兹勒乌",
		TitleCode: "b_tazlau",
	}
}
