package rogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约尼格达登JonegardenBarony struct {
	feud.BaseBarony
}

var BJonegarden约尼格达登 feud.Barony = &约尼格达登JonegardenBarony{}

func init() {
    f := BJonegarden约尼格达登.(*约尼格达登JonegardenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jonegarden",
		TitleName: "约尼格达登",
		TitleCode: "b_jonegarden",
	}
}
