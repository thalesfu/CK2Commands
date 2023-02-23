package braunschweig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔德斯海姆HildesheimBarony struct {
	feud.BaseBarony
}

var BHildesheim希尔德斯海姆 feud.Barony = &希尔德斯海姆HildesheimBarony{}

func init() {
	f := BHildesheim希尔德斯海姆.(*希尔德斯海姆HildesheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hildesheim",
		TitleName: "希尔德斯海姆",
		TitleCode: "b_hildesheim",
	}
}
