package cholamandalam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那伽钵亶那NagapattinamBarony struct {
	feud.BaseBarony
}

var BNagapattinam那伽钵亶那 feud.Barony = &那伽钵亶那NagapattinamBarony{}

func init() {
    f := BNagapattinam那伽钵亶那.(*那伽钵亶那NagapattinamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagapattinam",
		TitleName: "那伽钵亶那",
		TitleCode: "b_nagapattinam",
	}
}
