package theodosiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔图姆TortumBarony struct {
	feud.BaseBarony
}

var BTortum托尔图姆 feud.Barony = &托尔图姆TortumBarony{}

func init() {
    f := BTortum托尔图姆.(*托尔图姆TortumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tortum",
		TitleName: "托尔图姆",
		TitleCode: "b_tortum",
	}
}
