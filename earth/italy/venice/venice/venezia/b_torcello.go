package venezia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔切洛TorcelloBarony struct {
	feud.BaseBarony
}

var BTorcello托尔切洛 feud.Barony = &托尔切洛TorcelloBarony{}

func init() {
	f := BTorcello托尔切洛.(*托尔切洛TorcelloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torcello",
		TitleName: "托尔切洛",
		TitleCode: "b_torcello",
	}
}
