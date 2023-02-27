package lesser_poland

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/lesser_poland/krakowskie"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/lesser_poland/lubelska"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/lesser_poland/sacz"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/lesser_poland/sandomierskie"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/lesser_poland/stezycka"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Lesser_polandDuke interface {
    feud.Duke
    CKrakowskie克拉科夫() 	krakowskie.KrakowskieCounty
    CLubelska卢布林() 	lubelska.LubelskaCounty
    CSacz松奇() 	sacz.SaczCounty
    CSandomierskie桑多梅日() 	sandomierskie.SandomierskieCounty
    CStezycka斯滕日察() 	stezycka.StezyckaCounty
}

type 小波兰Lesser_polandDuke struct {
	feud.BaseDuke
	克拉科夫Krakowskie 	krakowskie.KrakowskieCounty
	卢布林Lubelska 	lubelska.LubelskaCounty
	松奇Sacz 	sacz.SaczCounty
	桑多梅日Sandomierskie 	sandomierskie.SandomierskieCounty
	斯滕日察Stezycka 	stezycka.StezyckaCounty
}

func (d *小波兰Lesser_polandDuke) CKrakowskie克拉科夫() krakowskie.KrakowskieCounty {
	return d.克拉科夫Krakowskie
}
    
func (d *小波兰Lesser_polandDuke) CLubelska卢布林() lubelska.LubelskaCounty {
	return d.卢布林Lubelska
}
    
func (d *小波兰Lesser_polandDuke) CSacz松奇() sacz.SaczCounty {
	return d.松奇Sacz
}
    
func (d *小波兰Lesser_polandDuke) CSandomierskie桑多梅日() sandomierskie.SandomierskieCounty {
	return d.桑多梅日Sandomierskie
}
    
func (d *小波兰Lesser_polandDuke) CStezycka斯滕日察() stezycka.StezyckaCounty {
	return d.斯滕日察Stezycka
}
    
var DLesser_poland小波兰 Lesser_polandDuke = &小波兰Lesser_polandDuke{}

func init() {
	f := DLesser_poland小波兰.(*小波兰Lesser_polandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lesser_poland",
		TitleName: "小波兰",
		TitleCode: "d_lesser_poland",
		Counties:  map[string]feud.County{},
	}

	f.克拉科夫Krakowskie = krakowskie.CKrakowskie克拉科夫
	f.克拉科夫Krakowskie.SetParent(f)
	
	f.卢布林Lubelska = lubelska.CLubelska卢布林
	f.卢布林Lubelska.SetParent(f)
	
	f.松奇Sacz = sacz.CSacz松奇
	f.松奇Sacz.SetParent(f)
	
	f.桑多梅日Sandomierskie = sandomierskie.CSandomierskie桑多梅日
	f.桑多梅日Sandomierskie.SetParent(f)
	
	f.斯滕日察Stezycka = stezycka.CStezycka斯滕日察
	f.斯滕日察Stezycka.SetParent(f)
	
}
