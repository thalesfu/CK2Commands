package pratishthana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃利施旃陀罗伽荼HarishchandragadBarony struct {
	feud.BaseBarony
}

var BHarishchandragad诃利施旃陀罗伽荼 feud.Barony = &诃利施旃陀罗伽荼HarishchandragadBarony{}

func init() {
	f := BHarishchandragad诃利施旃陀罗伽荼.(*诃利施旃陀罗伽荼HarishchandragadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harishchandragad",
		TitleName: "诃利施旃陀罗伽荼",
		TitleCode: "b_harishchandragad",
	}
}
