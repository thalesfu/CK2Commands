package ryn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙加泰ShagatayBarony struct {
	feud.BaseBarony
}

var BShagatay沙加泰 feud.Barony = &沙加泰ShagatayBarony{}

func init() {
    f := BShagatay沙加泰.(*沙加泰ShagatayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shagatay",
		TitleName: "沙加泰",
		TitleCode: "b_shagatay",
	}
}
