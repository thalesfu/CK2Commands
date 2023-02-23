package braunschweig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘德斯海姆GandersheimBarony struct {
	feud.BaseBarony
}

var BGandersheim甘德斯海姆 feud.Barony = &甘德斯海姆GandersheimBarony{}

func init() {
	f := BGandersheim甘德斯海姆.(*甘德斯海姆GandersheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gandersheim",
		TitleName: "甘德斯海姆",
		TitleCode: "b_gandersheim",
	}
}
