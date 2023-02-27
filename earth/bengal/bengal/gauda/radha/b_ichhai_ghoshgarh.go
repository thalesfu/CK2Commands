package radha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊车耶瞿沙姞利呬Ichhai_ghoshgarhBarony struct {
	feud.BaseBarony
}

var BIchhai_ghoshgarh伊车耶瞿沙姞利呬 feud.Barony = &伊车耶瞿沙姞利呬Ichhai_ghoshgarhBarony{}

func init() {
    f := BIchhai_ghoshgarh伊车耶瞿沙姞利呬.(*伊车耶瞿沙姞利呬Ichhai_ghoshgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ichhai_ghoshgarh",
		TitleName: "伊车耶瞿沙姞利呬",
		TitleCode: "b_ichhai_ghoshgarh",
	}
}
