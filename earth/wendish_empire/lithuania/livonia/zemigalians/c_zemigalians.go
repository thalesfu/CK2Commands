package zemigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZemigaliansCounty interface {
    feud.County
    BDobe多贝() 	feud.Barony
    BDobele多贝莱() 	feud.Barony
    BJelgawa叶尔加瓦() 	feud.Barony
    BMezotne梅若特内() 	feud.Barony
    BRakte拉克塔() 	feud.Barony
    BSparnene斯帕尔内内() 	feud.Barony
    BTervete泰尔韦泰() 	feud.Barony
}

type 泽姆加莱ZemigaliansCounty struct {
	feud.BaseCounty
	多贝Dobe 	feud.Barony
	多贝莱Dobele 	feud.Barony
	叶尔加瓦Jelgawa 	feud.Barony
	梅若特内Mezotne 	feud.Barony
	拉克塔Rakte 	feud.Barony
	斯帕尔内内Sparnene 	feud.Barony
	泰尔韦泰Tervete 	feud.Barony
}

func (c *泽姆加莱ZemigaliansCounty) BDobe多贝() feud.Barony {
	return c.多贝Dobe
}
    
func (c *泽姆加莱ZemigaliansCounty) BDobele多贝莱() feud.Barony {
	return c.多贝莱Dobele
}
    
func (c *泽姆加莱ZemigaliansCounty) BJelgawa叶尔加瓦() feud.Barony {
	return c.叶尔加瓦Jelgawa
}
    
func (c *泽姆加莱ZemigaliansCounty) BMezotne梅若特内() feud.Barony {
	return c.梅若特内Mezotne
}
    
func (c *泽姆加莱ZemigaliansCounty) BRakte拉克塔() feud.Barony {
	return c.拉克塔Rakte
}
    
func (c *泽姆加莱ZemigaliansCounty) BSparnene斯帕尔内内() feud.Barony {
	return c.斯帕尔内内Sparnene
}
    
func (c *泽姆加莱ZemigaliansCounty) BTervete泰尔韦泰() feud.Barony {
	return c.泰尔韦泰Tervete
}
    
var CZemigalians泽姆加莱 ZemigaliansCounty = &泽姆加莱ZemigaliansCounty{}

func init() {
	f := CZemigalians泽姆加莱.(*泽姆加莱ZemigaliansCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "374",
		Title:     "zemigalians",
		TitleName: "泽姆加莱",
		TitleCode: "c_zemigalians",
		Baronies:  map[string]feud.Barony{},
	}

	f.多贝Dobe = BDobe多贝
	f.多贝Dobe.SetParent(f)
	
	f.多贝莱Dobele = BDobele多贝莱
	f.多贝莱Dobele.SetParent(f)
	
	f.叶尔加瓦Jelgawa = BJelgawa叶尔加瓦
	f.叶尔加瓦Jelgawa.SetParent(f)
	
	f.梅若特内Mezotne = BMezotne梅若特内
	f.梅若特内Mezotne.SetParent(f)
	
	f.拉克塔Rakte = BRakte拉克塔
	f.拉克塔Rakte.SetParent(f)
	
	f.斯帕尔内内Sparnene = BSparnene斯帕尔内内
	f.斯帕尔内内Sparnene.SetParent(f)
	
	f.泰尔韦泰Tervete = BTervete泰尔韦泰
	f.泰尔韦泰Tervete.SetParent(f)
	
}
