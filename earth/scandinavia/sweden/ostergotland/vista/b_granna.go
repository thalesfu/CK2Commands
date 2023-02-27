package vista

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格伦纳GrannaBarony struct {
	feud.BaseBarony
}

var BGranna格伦纳 feud.Barony = &格伦纳GrannaBarony{}

func init() {
    f := BGranna格伦纳.(*格伦纳GrannaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "granna",
		TitleName: "格伦纳",
		TitleCode: "b_granna",
	}
}
