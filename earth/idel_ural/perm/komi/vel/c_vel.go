package vel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VelCounty interface {
    feud.County
    BAyuva阿尤瓦() 	feud.Barony
    BDutovo杜托沃() 	feud.Barony
    BKyrta克尔塔() 	feud.Barony
    BLemtybozh列姆特博日() 	feud.Barony
    BLemyu列米尤() 	feud.Barony
    BVel韦尔() 	feud.Barony
    BVuktyl_vel武克特尔() 	feud.Barony
}

type 韦尔VelCounty struct {
	feud.BaseCounty
	阿尤瓦Ayuva 	feud.Barony
	杜托沃Dutovo 	feud.Barony
	克尔塔Kyrta 	feud.Barony
	列姆特博日Lemtybozh 	feud.Barony
	列米尤Lemyu 	feud.Barony
	韦尔Vel 	feud.Barony
	武克特尔Vuktyl_vel 	feud.Barony
}

func (c *韦尔VelCounty) BAyuva阿尤瓦() feud.Barony {
	return c.阿尤瓦Ayuva
}
    
func (c *韦尔VelCounty) BDutovo杜托沃() feud.Barony {
	return c.杜托沃Dutovo
}
    
func (c *韦尔VelCounty) BKyrta克尔塔() feud.Barony {
	return c.克尔塔Kyrta
}
    
func (c *韦尔VelCounty) BLemtybozh列姆特博日() feud.Barony {
	return c.列姆特博日Lemtybozh
}
    
func (c *韦尔VelCounty) BLemyu列米尤() feud.Barony {
	return c.列米尤Lemyu
}
    
func (c *韦尔VelCounty) BVel韦尔() feud.Barony {
	return c.韦尔Vel
}
    
func (c *韦尔VelCounty) BVuktyl_vel武克特尔() feud.Barony {
	return c.武克特尔Vuktyl_vel
}
    
var CVel韦尔 VelCounty = &韦尔VelCounty{}

func init() {
	f := CVel韦尔.(*韦尔VelCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1836",
		Title:     "vel",
		TitleName: "韦尔",
		TitleCode: "c_vel",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尤瓦Ayuva = BAyuva阿尤瓦
	f.阿尤瓦Ayuva.SetParent(f)
	
	f.杜托沃Dutovo = BDutovo杜托沃
	f.杜托沃Dutovo.SetParent(f)
	
	f.克尔塔Kyrta = BKyrta克尔塔
	f.克尔塔Kyrta.SetParent(f)
	
	f.列姆特博日Lemtybozh = BLemtybozh列姆特博日
	f.列姆特博日Lemtybozh.SetParent(f)
	
	f.列米尤Lemyu = BLemyu列米尤
	f.列米尤Lemyu.SetParent(f)
	
	f.韦尔Vel = BVel韦尔
	f.韦尔Vel.SetParent(f)
	
	f.武克特尔Vuktyl_vel = BVuktyl_vel武克特尔
	f.武克特尔Vuktyl_vel.SetParent(f)
	
}
