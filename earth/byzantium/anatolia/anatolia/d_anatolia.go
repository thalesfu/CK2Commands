package anatolia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/anatolia/akroinon"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/anatolia/amorion"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/anatolia/ikonion"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnatoliaDuke interface {
    feud.Duke
    CAkroinon阿克罗伊农() 	akroinon.AkroinonCounty
    CAmorion阿摩里翁() 	amorion.AmorionCounty
    CIkonion以哥念() 	ikonion.IkonionCounty
}

type 安纳托利亚AnatoliaDuke struct {
	feud.BaseDuke
	阿克罗伊农Akroinon 	akroinon.AkroinonCounty
	阿摩里翁Amorion 	amorion.AmorionCounty
	以哥念Ikonion 	ikonion.IkonionCounty
}

func (d *安纳托利亚AnatoliaDuke) CAkroinon阿克罗伊农() akroinon.AkroinonCounty {
	return d.阿克罗伊农Akroinon
}
    
func (d *安纳托利亚AnatoliaDuke) CAmorion阿摩里翁() amorion.AmorionCounty {
	return d.阿摩里翁Amorion
}
    
func (d *安纳托利亚AnatoliaDuke) CIkonion以哥念() ikonion.IkonionCounty {
	return d.以哥念Ikonion
}
    
var DAnatolia安纳托利亚 AnatoliaDuke = &安纳托利亚AnatoliaDuke{}

func init() {
	f := DAnatolia安纳托利亚.(*安纳托利亚AnatoliaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "anatolia",
		TitleName: "安纳托利亚",
		TitleCode: "d_anatolia",
		Counties:  map[string]feud.County{},
	}

	f.阿克罗伊农Akroinon = akroinon.CAkroinon阿克罗伊农
	f.阿克罗伊农Akroinon.SetParent(f)
	
	f.阿摩里翁Amorion = amorion.CAmorion阿摩里翁
	f.阿摩里翁Amorion.SetParent(f)
	
	f.以哥念Ikonion = ikonion.CIkonion以哥念
	f.以哥念Ikonion.SetParent(f)
	
}
