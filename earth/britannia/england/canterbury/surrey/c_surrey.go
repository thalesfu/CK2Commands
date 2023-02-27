package surrey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SurreyCounty interface {
    feud.County
    BChertsey彻特西() 	feud.Barony
    BCroydon克罗伊登() 	feud.Barony
    BFarnham法纳姆() 	feud.Barony
    BGuildford吉尔福德() 	feud.Barony
    BLambeth朗伯斯() 	feud.Barony
    BSouthwark萨瑟克() 	feud.Barony
    BWaverley韦弗利() 	feud.Barony
    BWoking沃金() 	feud.Barony
}

type 萨里SurreyCounty struct {
	feud.BaseCounty
	彻特西Chertsey 	feud.Barony
	克罗伊登Croydon 	feud.Barony
	法纳姆Farnham 	feud.Barony
	吉尔福德Guildford 	feud.Barony
	朗伯斯Lambeth 	feud.Barony
	萨瑟克Southwark 	feud.Barony
	韦弗利Waverley 	feud.Barony
	沃金Woking 	feud.Barony
}

func (c *萨里SurreyCounty) BChertsey彻特西() feud.Barony {
	return c.彻特西Chertsey
}
    
func (c *萨里SurreyCounty) BCroydon克罗伊登() feud.Barony {
	return c.克罗伊登Croydon
}
    
func (c *萨里SurreyCounty) BFarnham法纳姆() feud.Barony {
	return c.法纳姆Farnham
}
    
func (c *萨里SurreyCounty) BGuildford吉尔福德() feud.Barony {
	return c.吉尔福德Guildford
}
    
func (c *萨里SurreyCounty) BLambeth朗伯斯() feud.Barony {
	return c.朗伯斯Lambeth
}
    
func (c *萨里SurreyCounty) BSouthwark萨瑟克() feud.Barony {
	return c.萨瑟克Southwark
}
    
func (c *萨里SurreyCounty) BWaverley韦弗利() feud.Barony {
	return c.韦弗利Waverley
}
    
func (c *萨里SurreyCounty) BWoking沃金() feud.Barony {
	return c.沃金Woking
}
    
var CSurrey萨里 SurreyCounty = &萨里SurreyCounty{}

func init() {
	f := CSurrey萨里.(*萨里SurreyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "24",
		Title:     "surrey",
		TitleName: "萨里",
		TitleCode: "c_surrey",
		Baronies:  map[string]feud.Barony{},
	}

	f.彻特西Chertsey = BChertsey彻特西
	f.彻特西Chertsey.SetParent(f)
	
	f.克罗伊登Croydon = BCroydon克罗伊登
	f.克罗伊登Croydon.SetParent(f)
	
	f.法纳姆Farnham = BFarnham法纳姆
	f.法纳姆Farnham.SetParent(f)
	
	f.吉尔福德Guildford = BGuildford吉尔福德
	f.吉尔福德Guildford.SetParent(f)
	
	f.朗伯斯Lambeth = BLambeth朗伯斯
	f.朗伯斯Lambeth.SetParent(f)
	
	f.萨瑟克Southwark = BSouthwark萨瑟克
	f.萨瑟克Southwark.SetParent(f)
	
	f.韦弗利Waverley = BWaverley韦弗利
	f.韦弗利Waverley.SetParent(f)
	
	f.沃金Woking = BWoking沃金
	f.沃金Woking.SetParent(f)
	
}
