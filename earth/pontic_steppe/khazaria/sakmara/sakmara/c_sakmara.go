package sakmara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SakmaraCounty interface {
    feud.County
    BBulanovo布拉诺沃() 	feud.Barony
    BIskra伊斯克拉() 	feud.Barony
    BKushkul库什库尔() 	feud.Barony
    BOrenburg奥伦堡() 	feud.Barony
    BPrudy普鲁迪() 	feud.Barony
    BSakmara萨克马拉() 	feud.Barony
    BZerklo泽尔克洛() 	feud.Barony
}

type 萨克马拉SakmaraCounty struct {
	feud.BaseCounty
	布拉诺沃Bulanovo 	feud.Barony
	伊斯克拉Iskra 	feud.Barony
	库什库尔Kushkul 	feud.Barony
	奥伦堡Orenburg 	feud.Barony
	普鲁迪Prudy 	feud.Barony
	萨克马拉Sakmara 	feud.Barony
	泽尔克洛Zerklo 	feud.Barony
}

func (c *萨克马拉SakmaraCounty) BBulanovo布拉诺沃() feud.Barony {
	return c.布拉诺沃Bulanovo
}
    
func (c *萨克马拉SakmaraCounty) BIskra伊斯克拉() feud.Barony {
	return c.伊斯克拉Iskra
}
    
func (c *萨克马拉SakmaraCounty) BKushkul库什库尔() feud.Barony {
	return c.库什库尔Kushkul
}
    
func (c *萨克马拉SakmaraCounty) BOrenburg奥伦堡() feud.Barony {
	return c.奥伦堡Orenburg
}
    
func (c *萨克马拉SakmaraCounty) BPrudy普鲁迪() feud.Barony {
	return c.普鲁迪Prudy
}
    
func (c *萨克马拉SakmaraCounty) BSakmara萨克马拉() feud.Barony {
	return c.萨克马拉Sakmara
}
    
func (c *萨克马拉SakmaraCounty) BZerklo泽尔克洛() feud.Barony {
	return c.泽尔克洛Zerklo
}
    
var CSakmara萨克马拉 SakmaraCounty = &萨克马拉SakmaraCounty{}

func init() {
	f := CSakmara萨克马拉.(*萨克马拉SakmaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1848",
		Title:     "sakmara",
		TitleName: "萨克马拉",
		TitleCode: "c_sakmara",
		Baronies:  map[string]feud.Barony{},
	}

	f.布拉诺沃Bulanovo = BBulanovo布拉诺沃
	f.布拉诺沃Bulanovo.SetParent(f)
	
	f.伊斯克拉Iskra = BIskra伊斯克拉
	f.伊斯克拉Iskra.SetParent(f)
	
	f.库什库尔Kushkul = BKushkul库什库尔
	f.库什库尔Kushkul.SetParent(f)
	
	f.奥伦堡Orenburg = BOrenburg奥伦堡
	f.奥伦堡Orenburg.SetParent(f)
	
	f.普鲁迪Prudy = BPrudy普鲁迪
	f.普鲁迪Prudy.SetParent(f)
	
	f.萨克马拉Sakmara = BSakmara萨克马拉
	f.萨克马拉Sakmara.SetParent(f)
	
	f.泽尔克洛Zerklo = BZerklo泽尔克洛
	f.泽尔克洛Zerklo.SetParent(f)
	
}
