package lithuanians

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/lithuanians/aukshayts"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/lithuanians/deltuva"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/lithuanians/nalsia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/lithuanians/yatvyagi"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/lithuanians/zhmud"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LithuaniansDuke interface {
    feud.Duke
    CAukshayts维尔纽斯() 	aukshayts.AukshaytsCounty
    CDeltuva德尔图瓦() 	deltuva.DeltuvaCounty
    CNalsia纳尔希亚() 	nalsia.NalsiaCounty
    CYatvyagi考纳斯() 	yatvyagi.YatvyagiCounty
    CZhmud萨莫吉提亚() 	zhmud.ZhmudCounty
}

type 立陶宛LithuaniansDuke struct {
	feud.BaseDuke
	维尔纽斯Aukshayts 	aukshayts.AukshaytsCounty
	德尔图瓦Deltuva 	deltuva.DeltuvaCounty
	纳尔希亚Nalsia 	nalsia.NalsiaCounty
	考纳斯Yatvyagi 	yatvyagi.YatvyagiCounty
	萨莫吉提亚Zhmud 	zhmud.ZhmudCounty
}

func (d *立陶宛LithuaniansDuke) CAukshayts维尔纽斯() aukshayts.AukshaytsCounty {
	return d.维尔纽斯Aukshayts
}
    
func (d *立陶宛LithuaniansDuke) CDeltuva德尔图瓦() deltuva.DeltuvaCounty {
	return d.德尔图瓦Deltuva
}
    
func (d *立陶宛LithuaniansDuke) CNalsia纳尔希亚() nalsia.NalsiaCounty {
	return d.纳尔希亚Nalsia
}
    
func (d *立陶宛LithuaniansDuke) CYatvyagi考纳斯() yatvyagi.YatvyagiCounty {
	return d.考纳斯Yatvyagi
}
    
func (d *立陶宛LithuaniansDuke) CZhmud萨莫吉提亚() zhmud.ZhmudCounty {
	return d.萨莫吉提亚Zhmud
}
    
var DLithuanians立陶宛 LithuaniansDuke = &立陶宛LithuaniansDuke{}

func init() {
	f := DLithuanians立陶宛.(*立陶宛LithuaniansDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lithuanians",
		TitleName: "立陶宛",
		TitleCode: "d_lithuanians",
		Counties:  map[string]feud.County{},
	}

	f.维尔纽斯Aukshayts = aukshayts.CAukshayts维尔纽斯
	f.维尔纽斯Aukshayts.SetParent(f)
	
	f.德尔图瓦Deltuva = deltuva.CDeltuva德尔图瓦
	f.德尔图瓦Deltuva.SetParent(f)
	
	f.纳尔希亚Nalsia = nalsia.CNalsia纳尔希亚
	f.纳尔希亚Nalsia.SetParent(f)
	
	f.考纳斯Yatvyagi = yatvyagi.CYatvyagi考纳斯
	f.考纳斯Yatvyagi.SetParent(f)
	
	f.萨莫吉提亚Zhmud = zhmud.CZhmud萨莫吉提亚
	f.萨莫吉提亚Zhmud.SetParent(f)
	
}
