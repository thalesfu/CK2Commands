package madurai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 愣那陀补愣RamnadhapuramBarony struct {
	feud.BaseBarony
}

var BRamnadhapuram愣那陀补愣 feud.Barony = &愣那陀补愣RamnadhapuramBarony{}

func init() {
    f := BRamnadhapuram愣那陀补愣.(*愣那陀补愣RamnadhapuramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramnadhapuram",
		TitleName: "愣那陀补愣",
		TitleCode: "b_ramnadhapuram",
	}
}
