package sarasvati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃陀罗ChandaraBarony struct {
	feud.BaseBarony
}

var BChandara旃陀罗 feud.Barony = &旃陀罗ChandaraBarony{}

func init() {
	f := BChandara旃陀罗.(*旃陀罗ChandaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chandara",
		TitleName: "旃陀罗",
		TitleCode: "b_chandara",
	}
}
