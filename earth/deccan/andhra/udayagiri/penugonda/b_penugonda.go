package penugonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卑奴郡荼PenugondaBarony struct {
	feud.BaseBarony
}

var BPenugonda卑奴郡荼 feud.Barony = &卑奴郡荼PenugondaBarony{}

func init() {
    f := BPenugonda卑奴郡荼.(*卑奴郡荼PenugondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "penugonda",
		TitleName: "卑奴郡荼",
		TitleCode: "b_penugonda",
	}
}
