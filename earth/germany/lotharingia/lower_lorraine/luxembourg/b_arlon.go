package luxembourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔隆ArlonBarony struct {
	feud.BaseBarony
}

var BArlon阿尔隆 feud.Barony = &阿尔隆ArlonBarony{}

func init() {
    f := BArlon阿尔隆.(*阿尔隆ArlonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arlon",
		TitleName: "阿尔隆",
		TitleCode: "b_arlon",
	}
}
