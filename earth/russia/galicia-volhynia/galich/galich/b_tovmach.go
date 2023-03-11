package galich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托夫马奇TovmachBarony struct {
	feud.BaseBarony
}

var BTovmach托夫马奇 feud.Barony = &托夫马奇TovmachBarony{}

func init() {
    f := BTovmach托夫马奇.(*托夫马奇TovmachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tovmach",
		TitleName: "托夫马奇",
		TitleCode: "b_tovmach",
	}
}
