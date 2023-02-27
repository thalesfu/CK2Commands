package nanded

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 势罗伽罗摩ShelgaonBarony struct {
	feud.BaseBarony
}

var BShelgaon势罗伽罗摩 feud.Barony = &势罗伽罗摩ShelgaonBarony{}

func init() {
    f := BShelgaon势罗伽罗摩.(*势罗伽罗摩ShelgaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shelgaon",
		TitleName: "势罗伽罗摩",
		TitleCode: "b_shelgaon",
	}
}
