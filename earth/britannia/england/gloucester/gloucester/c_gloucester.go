package gloucester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GloucesterCounty interface {
    feud.County
    BCheltenham切尔滕纳姆() 	feud.Barony
    BCirencester赛伦塞斯特() 	feud.Barony
    BGloucester格洛斯特() 	feud.Barony
    BHailes黑尔斯() 	feud.Barony
    BSudeley休德利() 	feud.Barony
    BTewkesbury蒂克斯伯里() 	feud.Barony
    BWinchcombe温什科姆() 	feud.Barony
}

type 格洛斯特GloucesterCounty struct {
	feud.BaseCounty
	切尔滕纳姆Cheltenham 	feud.Barony
	赛伦塞斯特Cirencester 	feud.Barony
	格洛斯特Gloucester 	feud.Barony
	黑尔斯Hailes 	feud.Barony
	休德利Sudeley 	feud.Barony
	蒂克斯伯里Tewkesbury 	feud.Barony
	温什科姆Winchcombe 	feud.Barony
}

func (c *格洛斯特GloucesterCounty) BCheltenham切尔滕纳姆() feud.Barony {
	return c.切尔滕纳姆Cheltenham
}
    
func (c *格洛斯特GloucesterCounty) BCirencester赛伦塞斯特() feud.Barony {
	return c.赛伦塞斯特Cirencester
}
    
func (c *格洛斯特GloucesterCounty) BGloucester格洛斯特() feud.Barony {
	return c.格洛斯特Gloucester
}
    
func (c *格洛斯特GloucesterCounty) BHailes黑尔斯() feud.Barony {
	return c.黑尔斯Hailes
}
    
func (c *格洛斯特GloucesterCounty) BSudeley休德利() feud.Barony {
	return c.休德利Sudeley
}
    
func (c *格洛斯特GloucesterCounty) BTewkesbury蒂克斯伯里() feud.Barony {
	return c.蒂克斯伯里Tewkesbury
}
    
func (c *格洛斯特GloucesterCounty) BWinchcombe温什科姆() feud.Barony {
	return c.温什科姆Winchcombe
}
    
var CGloucester格洛斯特 GloucesterCounty = &格洛斯特GloucesterCounty{}

func init() {
	f := CGloucester格洛斯特.(*格洛斯特GloucesterCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "21",
		Title:     "gloucester",
		TitleName: "格洛斯特",
		TitleCode: "c_gloucester",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔滕纳姆Cheltenham = BCheltenham切尔滕纳姆
	f.切尔滕纳姆Cheltenham.SetParent(f)
	
	f.赛伦塞斯特Cirencester = BCirencester赛伦塞斯特
	f.赛伦塞斯特Cirencester.SetParent(f)
	
	f.格洛斯特Gloucester = BGloucester格洛斯特
	f.格洛斯特Gloucester.SetParent(f)
	
	f.黑尔斯Hailes = BHailes黑尔斯
	f.黑尔斯Hailes.SetParent(f)
	
	f.休德利Sudeley = BSudeley休德利
	f.休德利Sudeley.SetParent(f)
	
	f.蒂克斯伯里Tewkesbury = BTewkesbury蒂克斯伯里
	f.蒂克斯伯里Tewkesbury.SetParent(f)
	
	f.温什科姆Winchcombe = BWinchcombe温什科姆
	f.温什科姆Winchcombe.SetParent(f)
	
}
