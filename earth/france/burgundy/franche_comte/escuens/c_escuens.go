package escuens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EscuensCounty interface {
    feud.County
    BBelin伯兰() 	feud.Barony
    BBracon布拉孔() 	feud.Barony
    BFrontenay弗龙特奈() 	feud.Barony
    BGuyon吉永() 	feud.Barony
    BNozeroy诺兹鲁瓦() 	feud.Barony
    BPin潘() 	feud.Barony
    BPoligny波利尼() 	feud.Barony
    BSaint_claude圣克洛德() 	feud.Barony
    BSalins萨兰() 	feud.Barony
}

type 埃斯库昂EscuensCounty struct {
	feud.BaseCounty
	伯兰Belin 	feud.Barony
	布拉孔Bracon 	feud.Barony
	弗龙特奈Frontenay 	feud.Barony
	吉永Guyon 	feud.Barony
	诺兹鲁瓦Nozeroy 	feud.Barony
	潘Pin 	feud.Barony
	波利尼Poligny 	feud.Barony
	圣克洛德Saint_claude 	feud.Barony
	萨兰Salins 	feud.Barony
}

func (c *埃斯库昂EscuensCounty) BBelin伯兰() feud.Barony {
	return c.伯兰Belin
}
    
func (c *埃斯库昂EscuensCounty) BBracon布拉孔() feud.Barony {
	return c.布拉孔Bracon
}
    
func (c *埃斯库昂EscuensCounty) BFrontenay弗龙特奈() feud.Barony {
	return c.弗龙特奈Frontenay
}
    
func (c *埃斯库昂EscuensCounty) BGuyon吉永() feud.Barony {
	return c.吉永Guyon
}
    
func (c *埃斯库昂EscuensCounty) BNozeroy诺兹鲁瓦() feud.Barony {
	return c.诺兹鲁瓦Nozeroy
}
    
func (c *埃斯库昂EscuensCounty) BPin潘() feud.Barony {
	return c.潘Pin
}
    
func (c *埃斯库昂EscuensCounty) BPoligny波利尼() feud.Barony {
	return c.波利尼Poligny
}
    
func (c *埃斯库昂EscuensCounty) BSaint_claude圣克洛德() feud.Barony {
	return c.圣克洛德Saint_claude
}
    
func (c *埃斯库昂EscuensCounty) BSalins萨兰() feud.Barony {
	return c.萨兰Salins
}
    
var CEscuens埃斯库昂 EscuensCounty = &埃斯库昂EscuensCounty{}

func init() {
	f := CEscuens埃斯库昂.(*埃斯库昂EscuensCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1614",
		Title:     "escuens",
		TitleName: "埃斯库昂",
		TitleCode: "c_escuens",
		Baronies:  map[string]feud.Barony{},
	}

	f.伯兰Belin = BBelin伯兰
	f.伯兰Belin.SetParent(f)
	
	f.布拉孔Bracon = BBracon布拉孔
	f.布拉孔Bracon.SetParent(f)
	
	f.弗龙特奈Frontenay = BFrontenay弗龙特奈
	f.弗龙特奈Frontenay.SetParent(f)
	
	f.吉永Guyon = BGuyon吉永
	f.吉永Guyon.SetParent(f)
	
	f.诺兹鲁瓦Nozeroy = BNozeroy诺兹鲁瓦
	f.诺兹鲁瓦Nozeroy.SetParent(f)
	
	f.潘Pin = BPin潘
	f.潘Pin.SetParent(f)
	
	f.波利尼Poligny = BPoligny波利尼
	f.波利尼Poligny.SetParent(f)
	
	f.圣克洛德Saint_claude = BSaint_claude圣克洛德
	f.圣克洛德Saint_claude.SetParent(f)
	
	f.萨兰Salins = BSalins萨兰
	f.萨兰Salins.SetParent(f)
	
}
