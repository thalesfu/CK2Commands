package abauj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbaujCounty interface {
    feud.County
    BAbauj奥鲍乌伊() 	feud.Barony
    BKassa考绍() 	feud.Barony
    BSarospatak沙罗什保陶克() 	feud.Barony
    BSatoraljaujhely沙托劳尔尧乌伊海伊() 	feud.Barony
    BSzepsi塞普希() 	feud.Barony
    BSzikszo锡克索() 	feud.Barony
    BTokaj托考伊() 	feud.Barony
    BTurna图尔尼亚() 	feud.Barony
}

type 奥鲍乌伊AbaujCounty struct {
	feud.BaseCounty
	奥鲍乌伊Abauj 	feud.Barony
	考绍Kassa 	feud.Barony
	沙罗什保陶克Sarospatak 	feud.Barony
	沙托劳尔尧乌伊海伊Satoraljaujhely 	feud.Barony
	塞普希Szepsi 	feud.Barony
	锡克索Szikszo 	feud.Barony
	托考伊Tokaj 	feud.Barony
	图尔尼亚Turna 	feud.Barony
}

func (c *奥鲍乌伊AbaujCounty) BAbauj奥鲍乌伊() feud.Barony {
	return c.奥鲍乌伊Abauj
}
    
func (c *奥鲍乌伊AbaujCounty) BKassa考绍() feud.Barony {
	return c.考绍Kassa
}
    
func (c *奥鲍乌伊AbaujCounty) BSarospatak沙罗什保陶克() feud.Barony {
	return c.沙罗什保陶克Sarospatak
}
    
func (c *奥鲍乌伊AbaujCounty) BSatoraljaujhely沙托劳尔尧乌伊海伊() feud.Barony {
	return c.沙托劳尔尧乌伊海伊Satoraljaujhely
}
    
func (c *奥鲍乌伊AbaujCounty) BSzepsi塞普希() feud.Barony {
	return c.塞普希Szepsi
}
    
func (c *奥鲍乌伊AbaujCounty) BSzikszo锡克索() feud.Barony {
	return c.锡克索Szikszo
}
    
func (c *奥鲍乌伊AbaujCounty) BTokaj托考伊() feud.Barony {
	return c.托考伊Tokaj
}
    
func (c *奥鲍乌伊AbaujCounty) BTurna图尔尼亚() feud.Barony {
	return c.图尔尼亚Turna
}
    
var CAbauj奥鲍乌伊 AbaujCounty = &奥鲍乌伊AbaujCounty{}

func init() {
	f := CAbauj奥鲍乌伊.(*奥鲍乌伊AbaujCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "538",
		Title:     "abauj",
		TitleName: "奥鲍乌伊",
		TitleCode: "c_abauj",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥鲍乌伊Abauj = BAbauj奥鲍乌伊
	f.奥鲍乌伊Abauj.SetParent(f)
	
	f.考绍Kassa = BKassa考绍
	f.考绍Kassa.SetParent(f)
	
	f.沙罗什保陶克Sarospatak = BSarospatak沙罗什保陶克
	f.沙罗什保陶克Sarospatak.SetParent(f)
	
	f.沙托劳尔尧乌伊海伊Satoraljaujhely = BSatoraljaujhely沙托劳尔尧乌伊海伊
	f.沙托劳尔尧乌伊海伊Satoraljaujhely.SetParent(f)
	
	f.塞普希Szepsi = BSzepsi塞普希
	f.塞普希Szepsi.SetParent(f)
	
	f.锡克索Szikszo = BSzikszo锡克索
	f.锡克索Szikszo.SetParent(f)
	
	f.托考伊Tokaj = BTokaj托考伊
	f.托考伊Tokaj.SetParent(f)
	
	f.图尔尼亚Turna = BTurna图尔尼亚
	f.图尔尼亚Turna.SetParent(f)
	
}
