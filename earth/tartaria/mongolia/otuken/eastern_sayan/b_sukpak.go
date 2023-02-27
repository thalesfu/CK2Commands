package eastern_sayan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏克帕克SukpakBarony struct {
	feud.BaseBarony
}

var BSukpak苏克帕克 feud.Barony = &苏克帕克SukpakBarony{}

func init() {
    f := BSukpak苏克帕克.(*苏克帕克SukpakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sukpak",
		TitleName: "苏克帕克",
		TitleCode: "b_sukpak",
	}
}
