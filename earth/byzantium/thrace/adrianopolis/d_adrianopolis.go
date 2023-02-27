package adrianopolis

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/adrianopolis/adrianopolis"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/adrianopolis/maroneia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/adrianopolis/philippopolis"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/adrianopolis/traianopolis"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AdrianopolisDuke interface {
    feud.Duke
    CAdrianopolis哈德良波利斯() 	adrianopolis.AdrianopolisCounty
    CMaroneia马戎尼亚() 	maroneia.MaroneiaCounty
    CPhilippopolis腓利波波利斯() 	philippopolis.PhilippopolisCounty
    CTraianopolis图拉真波利斯() 	traianopolis.TraianopolisCounty
}

type 哈德良波利斯AdrianopolisDuke struct {
	feud.BaseDuke
	哈德良波利斯Adrianopolis 	adrianopolis.AdrianopolisCounty
	马戎尼亚Maroneia 	maroneia.MaroneiaCounty
	腓利波波利斯Philippopolis 	philippopolis.PhilippopolisCounty
	图拉真波利斯Traianopolis 	traianopolis.TraianopolisCounty
}

func (d *哈德良波利斯AdrianopolisDuke) CAdrianopolis哈德良波利斯() adrianopolis.AdrianopolisCounty {
	return d.哈德良波利斯Adrianopolis
}
    
func (d *哈德良波利斯AdrianopolisDuke) CMaroneia马戎尼亚() maroneia.MaroneiaCounty {
	return d.马戎尼亚Maroneia
}
    
func (d *哈德良波利斯AdrianopolisDuke) CPhilippopolis腓利波波利斯() philippopolis.PhilippopolisCounty {
	return d.腓利波波利斯Philippopolis
}
    
func (d *哈德良波利斯AdrianopolisDuke) CTraianopolis图拉真波利斯() traianopolis.TraianopolisCounty {
	return d.图拉真波利斯Traianopolis
}
    
var DAdrianopolis哈德良波利斯 AdrianopolisDuke = &哈德良波利斯AdrianopolisDuke{}

func init() {
	f := DAdrianopolis哈德良波利斯.(*哈德良波利斯AdrianopolisDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "adrianopolis",
		TitleName: "哈德良波利斯",
		TitleCode: "d_adrianopolis",
		Counties:  map[string]feud.County{},
	}

	f.哈德良波利斯Adrianopolis = adrianopolis.CAdrianopolis哈德良波利斯
	f.哈德良波利斯Adrianopolis.SetParent(f)
	
	f.马戎尼亚Maroneia = maroneia.CMaroneia马戎尼亚
	f.马戎尼亚Maroneia.SetParent(f)
	
	f.腓利波波利斯Philippopolis = philippopolis.CPhilippopolis腓利波波利斯
	f.腓利波波利斯Philippopolis.SetParent(f)
	
	f.图拉真波利斯Traianopolis = traianopolis.CTraianopolis图拉真波利斯
	f.图拉真波利斯Traianopolis.SetParent(f)
	
}
