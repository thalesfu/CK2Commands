package druz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DruzCounty interface {
    feud.County
    BAwas阿巴塞() 	feud.Barony
    BBusra布斯拉() 	feud.Barony
    BDibin迪比因() 	feud.Barony
    BGhariyah古哈利耶() 	feud.Barony
    BSalkhard塞勒海德() 	feud.Barony
    BSamj萨姆吉() 	feud.Barony
    BShannireh塞涅里赫() 	feud.Barony
    BThoula特霍拉() 	feud.Barony
}

type 德鲁兹DruzCounty struct {
	feud.BaseCounty
	阿巴塞Awas 	feud.Barony
	布斯拉Busra 	feud.Barony
	迪比因Dibin 	feud.Barony
	古哈利耶Ghariyah 	feud.Barony
	塞勒海德Salkhard 	feud.Barony
	萨姆吉Samj 	feud.Barony
	塞涅里赫Shannireh 	feud.Barony
	特霍拉Thoula 	feud.Barony
}

func (c *德鲁兹DruzCounty) BAwas阿巴塞() feud.Barony {
	return c.阿巴塞Awas
}
    
func (c *德鲁兹DruzCounty) BBusra布斯拉() feud.Barony {
	return c.布斯拉Busra
}
    
func (c *德鲁兹DruzCounty) BDibin迪比因() feud.Barony {
	return c.迪比因Dibin
}
    
func (c *德鲁兹DruzCounty) BGhariyah古哈利耶() feud.Barony {
	return c.古哈利耶Ghariyah
}
    
func (c *德鲁兹DruzCounty) BSalkhard塞勒海德() feud.Barony {
	return c.塞勒海德Salkhard
}
    
func (c *德鲁兹DruzCounty) BSamj萨姆吉() feud.Barony {
	return c.萨姆吉Samj
}
    
func (c *德鲁兹DruzCounty) BShannireh塞涅里赫() feud.Barony {
	return c.塞涅里赫Shannireh
}
    
func (c *德鲁兹DruzCounty) BThoula特霍拉() feud.Barony {
	return c.特霍拉Thoula
}
    
var CDruz德鲁兹 DruzCounty = &德鲁兹DruzCounty{}

func init() {
	f := CDruz德鲁兹.(*德鲁兹DruzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "713",
		Title:     "druz",
		TitleName: "德鲁兹",
		TitleCode: "c_druz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴塞Awas = BAwas阿巴塞
	f.阿巴塞Awas.SetParent(f)
	
	f.布斯拉Busra = BBusra布斯拉
	f.布斯拉Busra.SetParent(f)
	
	f.迪比因Dibin = BDibin迪比因
	f.迪比因Dibin.SetParent(f)
	
	f.古哈利耶Ghariyah = BGhariyah古哈利耶
	f.古哈利耶Ghariyah.SetParent(f)
	
	f.塞勒海德Salkhard = BSalkhard塞勒海德
	f.塞勒海德Salkhard.SetParent(f)
	
	f.萨姆吉Samj = BSamj萨姆吉
	f.萨姆吉Samj.SetParent(f)
	
	f.塞涅里赫Shannireh = BShannireh塞涅里赫
	f.塞涅里赫Shannireh.SetParent(f)
	
	f.特霍拉Thoula = BThoula特霍拉
	f.特霍拉Thoula.SetParent(f)
	
}
