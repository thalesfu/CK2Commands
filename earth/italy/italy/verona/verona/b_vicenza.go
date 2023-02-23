package verona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维琴察VicenzaBarony struct {
	feud.BaseBarony
}

var BVicenza维琴察 feud.Barony = &维琴察VicenzaBarony{}

func init() {
	f := BVicenza维琴察.(*维琴察VicenzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vicenza",
		TitleName: "维琴察",
		TitleCode: "b_vicenza",
	}
}
