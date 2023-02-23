package schwyz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔特多夫AltdorfBarony struct {
	feud.BaseBarony
}

var BAltdorf阿尔特多夫 feud.Barony = &阿尔特多夫AltdorfBarony{}

func init() {
	f := BAltdorf阿尔特多夫.(*阿尔特多夫AltdorfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altdorf",
		TitleName: "阿尔特多夫",
		TitleCode: "b_altdorf",
	}
}
