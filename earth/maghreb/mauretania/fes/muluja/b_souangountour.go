package muluja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏旺贡图尔SouangountourBarony struct {
	feud.BaseBarony
}

var BSouangountour苏旺贡图尔 feud.Barony = &苏旺贡图尔SouangountourBarony{}

func init() {
    f := BSouangountour苏旺贡图尔.(*苏旺贡图尔SouangountourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "souangountour",
		TitleName: "苏旺贡图尔",
		TitleCode: "b_souangountour",
	}
}
