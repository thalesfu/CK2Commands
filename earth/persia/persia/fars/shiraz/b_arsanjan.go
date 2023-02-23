package shiraz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿桑詹ArsanjanBarony struct {
	feud.BaseBarony
}

var BArsanjan阿桑詹 feud.Barony = &阿桑詹ArsanjanBarony{}

func init() {
	f := BArsanjan阿桑詹.(*阿桑詹ArsanjanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arsanjan",
		TitleName: "阿桑詹",
		TitleCode: "b_arsanjan",
	}
}
