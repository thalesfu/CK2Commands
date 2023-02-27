package zamora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZamoraCounty interface {
    feud.County
    BBenavente贝纳文特() 	feud.Barony
    BCorrales科拉莱斯() 	feud.Barony
    BFermoselle费尔莫塞列() 	feud.Barony
    BFuentesauco丰特萨乌科() 	feud.Barony
    BPolvorosa波尔沃罗萨() 	feud.Barony
    BSanabria萨纳夫里亚() 	feud.Barony
    BToro托罗() 	feud.Barony
    BZamora萨莫拉() 	feud.Barony
}

type 萨莫拉ZamoraCounty struct {
	feud.BaseCounty
	贝纳文特Benavente 	feud.Barony
	科拉莱斯Corrales 	feud.Barony
	费尔莫塞列Fermoselle 	feud.Barony
	丰特萨乌科Fuentesauco 	feud.Barony
	波尔沃罗萨Polvorosa 	feud.Barony
	萨纳夫里亚Sanabria 	feud.Barony
	托罗Toro 	feud.Barony
	萨莫拉Zamora 	feud.Barony
}

func (c *萨莫拉ZamoraCounty) BBenavente贝纳文特() feud.Barony {
	return c.贝纳文特Benavente
}
    
func (c *萨莫拉ZamoraCounty) BCorrales科拉莱斯() feud.Barony {
	return c.科拉莱斯Corrales
}
    
func (c *萨莫拉ZamoraCounty) BFermoselle费尔莫塞列() feud.Barony {
	return c.费尔莫塞列Fermoselle
}
    
func (c *萨莫拉ZamoraCounty) BFuentesauco丰特萨乌科() feud.Barony {
	return c.丰特萨乌科Fuentesauco
}
    
func (c *萨莫拉ZamoraCounty) BPolvorosa波尔沃罗萨() feud.Barony {
	return c.波尔沃罗萨Polvorosa
}
    
func (c *萨莫拉ZamoraCounty) BSanabria萨纳夫里亚() feud.Barony {
	return c.萨纳夫里亚Sanabria
}
    
func (c *萨莫拉ZamoraCounty) BToro托罗() feud.Barony {
	return c.托罗Toro
}
    
func (c *萨莫拉ZamoraCounty) BZamora萨莫拉() feud.Barony {
	return c.萨莫拉Zamora
}
    
var CZamora萨莫拉 ZamoraCounty = &萨莫拉ZamoraCounty{}

func init() {
	f := CZamora萨莫拉.(*萨莫拉ZamoraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "191",
		Title:     "zamora",
		TitleName: "萨莫拉",
		TitleCode: "c_zamora",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝纳文特Benavente = BBenavente贝纳文特
	f.贝纳文特Benavente.SetParent(f)
	
	f.科拉莱斯Corrales = BCorrales科拉莱斯
	f.科拉莱斯Corrales.SetParent(f)
	
	f.费尔莫塞列Fermoselle = BFermoselle费尔莫塞列
	f.费尔莫塞列Fermoselle.SetParent(f)
	
	f.丰特萨乌科Fuentesauco = BFuentesauco丰特萨乌科
	f.丰特萨乌科Fuentesauco.SetParent(f)
	
	f.波尔沃罗萨Polvorosa = BPolvorosa波尔沃罗萨
	f.波尔沃罗萨Polvorosa.SetParent(f)
	
	f.萨纳夫里亚Sanabria = BSanabria萨纳夫里亚
	f.萨纳夫里亚Sanabria.SetParent(f)
	
	f.托罗Toro = BToro托罗
	f.托罗Toro.SetParent(f)
	
	f.萨莫拉Zamora = BZamora萨莫拉
	f.萨莫拉Zamora.SetParent(f)
	
}
