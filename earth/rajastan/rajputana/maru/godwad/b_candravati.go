package godwad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃陀罗伐底CandravatiBarony struct {
	feud.BaseBarony
}

var BCandravati旃陀罗伐底 feud.Barony = &旃陀罗伐底CandravatiBarony{}

func init() {
	f := BCandravati旃陀罗伐底.(*旃陀罗伐底CandravatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "candravati",
		TitleName: "旃陀罗伐底",
		TitleCode: "b_candravati",
	}
}
