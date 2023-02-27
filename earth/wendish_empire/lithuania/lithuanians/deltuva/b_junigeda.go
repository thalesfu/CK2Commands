package deltuva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤尼格达JunigedaBarony struct {
	feud.BaseBarony
}

var BJunigeda尤尼格达 feud.Barony = &尤尼格达JunigedaBarony{}

func init() {
    f := BJunigeda尤尼格达.(*尤尼格达JunigedaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "junigeda",
		TitleName: "尤尼格达",
		TitleCode: "b_junigeda",
	}
}
