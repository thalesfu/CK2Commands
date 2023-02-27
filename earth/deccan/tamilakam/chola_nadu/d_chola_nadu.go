package chola_nadu

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/chola_nadu/cholamandalam"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam/chola_nadu/tagadur"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Chola_naduDuke interface {
    feud.Duke
    CCholamandalam朱罗曼荼罗() 	cholamandalam.CholamandalamCounty
    CTagadur多伽头罗() 	tagadur.TagadurCounty
}

type 朱罗拿柱Chola_naduDuke struct {
	feud.BaseDuke
	朱罗曼荼罗Cholamandalam 	cholamandalam.CholamandalamCounty
	多伽头罗Tagadur 	tagadur.TagadurCounty
}

func (d *朱罗拿柱Chola_naduDuke) CCholamandalam朱罗曼荼罗() cholamandalam.CholamandalamCounty {
	return d.朱罗曼荼罗Cholamandalam
}
    
func (d *朱罗拿柱Chola_naduDuke) CTagadur多伽头罗() tagadur.TagadurCounty {
	return d.多伽头罗Tagadur
}
    
var DChola_nadu朱罗拿柱 Chola_naduDuke = &朱罗拿柱Chola_naduDuke{}

func init() {
	f := DChola_nadu朱罗拿柱.(*朱罗拿柱Chola_naduDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "chola_nadu",
		TitleName: "朱罗拿柱",
		TitleCode: "d_chola_nadu",
		Counties:  map[string]feud.County{},
	}

	f.朱罗曼荼罗Cholamandalam = cholamandalam.CCholamandalam朱罗曼荼罗
	f.朱罗曼荼罗Cholamandalam.SetParent(f)
	
	f.多伽头罗Tagadur = tagadur.CTagadur多伽头罗
	f.多伽头罗Tagadur.SetParent(f)
	
}
