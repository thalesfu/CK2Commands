package gurgi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙吞SathonBarony struct {
	feud.BaseBarony
}

var BSathon沙吞 feud.Barony = &沙吞SathonBarony{}

func init() {
    f := BSathon沙吞.(*沙吞SathonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sathon",
		TitleName: "沙吞",
		TitleCode: "b_sathon",
	}
}
