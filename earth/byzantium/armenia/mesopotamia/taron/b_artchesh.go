package taron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿彻什ArtcheshBarony struct {
	feud.BaseBarony
}

var BArtchesh阿彻什 feud.Barony = &阿彻什ArtcheshBarony{}

func init() {
    f := BArtchesh阿彻什.(*阿彻什ArtcheshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "artchesh",
		TitleName: "阿彻什",
		TitleCode: "b_artchesh",
	}
}
