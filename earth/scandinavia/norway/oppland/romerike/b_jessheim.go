package romerike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰斯海姆JessheimBarony struct {
	feud.BaseBarony
}

var BJessheim杰斯海姆 feud.Barony = &杰斯海姆JessheimBarony{}

func init() {
	f := BJessheim杰斯海姆.(*杰斯海姆JessheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jessheim",
		TitleName: "杰斯海姆",
		TitleCode: "b_jessheim",
	}
}
