package sarasvata_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿那醯罗波吒迦AnahilapatakaBarony struct {
	feud.BaseBarony
}

var BAnahilapataka阿那醯罗波吒迦 feud.Barony = &阿那醯罗波吒迦AnahilapatakaBarony{}

func init() {
    f := BAnahilapataka阿那醯罗波吒迦.(*阿那醯罗波吒迦AnahilapatakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anahilapataka",
		TitleName: "阿那醯罗波吒迦",
		TitleCode: "b_anahilapataka",
	}
}
