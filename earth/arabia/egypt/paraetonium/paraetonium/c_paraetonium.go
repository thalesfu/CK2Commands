package paraetonium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ParaetoniumCounty interface {
    feud.County
    BAl_hadid哈迪德() 	feud.Barony
    BApis阿庇斯() 	feud.Barony
    BLeuce琉刻() 	feud.Barony
    BParaetonium帕莱托尼翁() 	feud.Barony
    BPhocusa福库萨() 	feud.Barony
    BSelinais塞利涅斯() 	feud.Barony
    BZygis卒吉斯() 	feud.Barony
}

type 帕莱托尼翁ParaetoniumCounty struct {
	feud.BaseCounty
	哈迪德Al_hadid 	feud.Barony
	阿庇斯Apis 	feud.Barony
	琉刻Leuce 	feud.Barony
	帕莱托尼翁Paraetonium 	feud.Barony
	福库萨Phocusa 	feud.Barony
	塞利涅斯Selinais 	feud.Barony
	卒吉斯Zygis 	feud.Barony
}

func (c *帕莱托尼翁ParaetoniumCounty) BAl_hadid哈迪德() feud.Barony {
	return c.哈迪德Al_hadid
}
    
func (c *帕莱托尼翁ParaetoniumCounty) BApis阿庇斯() feud.Barony {
	return c.阿庇斯Apis
}
    
func (c *帕莱托尼翁ParaetoniumCounty) BLeuce琉刻() feud.Barony {
	return c.琉刻Leuce
}
    
func (c *帕莱托尼翁ParaetoniumCounty) BParaetonium帕莱托尼翁() feud.Barony {
	return c.帕莱托尼翁Paraetonium
}
    
func (c *帕莱托尼翁ParaetoniumCounty) BPhocusa福库萨() feud.Barony {
	return c.福库萨Phocusa
}
    
func (c *帕莱托尼翁ParaetoniumCounty) BSelinais塞利涅斯() feud.Barony {
	return c.塞利涅斯Selinais
}
    
func (c *帕莱托尼翁ParaetoniumCounty) BZygis卒吉斯() feud.Barony {
	return c.卒吉斯Zygis
}
    
var CParaetonium帕莱托尼翁 ParaetoniumCounty = &帕莱托尼翁ParaetoniumCounty{}

func init() {
	f := CParaetonium帕莱托尼翁.(*帕莱托尼翁ParaetoniumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2005",
		Title:     "paraetonium",
		TitleName: "帕莱托尼翁",
		TitleCode: "c_paraetonium",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈迪德Al_hadid = BAl_hadid哈迪德
	f.哈迪德Al_hadid.SetParent(f)
	
	f.阿庇斯Apis = BApis阿庇斯
	f.阿庇斯Apis.SetParent(f)
	
	f.琉刻Leuce = BLeuce琉刻
	f.琉刻Leuce.SetParent(f)
	
	f.帕莱托尼翁Paraetonium = BParaetonium帕莱托尼翁
	f.帕莱托尼翁Paraetonium.SetParent(f)
	
	f.福库萨Phocusa = BPhocusa福库萨
	f.福库萨Phocusa.SetParent(f)
	
	f.塞利涅斯Selinais = BSelinais塞利涅斯
	f.塞利涅斯Selinais.SetParent(f)
	
	f.卒吉斯Zygis = BZygis卒吉斯
	f.卒吉斯Zygis.SetParent(f)
	
}
