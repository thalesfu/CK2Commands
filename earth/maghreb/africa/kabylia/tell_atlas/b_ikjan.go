package tell_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊克詹IkjanBarony struct {
	feud.BaseBarony
}

var BIkjan伊克詹 feud.Barony = &伊克詹IkjanBarony{}

func init() {
    f := BIkjan伊克詹.(*伊克詹IkjanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikjan",
		TitleName: "伊克詹",
		TitleCode: "b_ikjan",
	}
}
