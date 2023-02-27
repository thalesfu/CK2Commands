package kubera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KuberaCounty interface {
    feud.County
    BAxiyan阿悉言() 	feud.Barony
    BDazhou_fenchang大州分场() 	feud.Barony
    BGexianmiao葛仙庙() 	feud.Barony
    BKubera俱毗罗() 	feud.Barony
    BShitun石屯() 	feud.Barony
    BToqsu托克苏() 	feud.Barony
    BYangtaoshan杨桃山() 	feud.Barony
}

type 俱毗罗KuberaCounty struct {
	feud.BaseCounty
	阿悉言Axiyan 	feud.Barony
	大州分场Dazhou_fenchang 	feud.Barony
	葛仙庙Gexianmiao 	feud.Barony
	俱毗罗Kubera 	feud.Barony
	石屯Shitun 	feud.Barony
	托克苏Toqsu 	feud.Barony
	杨桃山Yangtaoshan 	feud.Barony
}

func (c *俱毗罗KuberaCounty) BAxiyan阿悉言() feud.Barony {
	return c.阿悉言Axiyan
}
    
func (c *俱毗罗KuberaCounty) BDazhou_fenchang大州分场() feud.Barony {
	return c.大州分场Dazhou_fenchang
}
    
func (c *俱毗罗KuberaCounty) BGexianmiao葛仙庙() feud.Barony {
	return c.葛仙庙Gexianmiao
}
    
func (c *俱毗罗KuberaCounty) BKubera俱毗罗() feud.Barony {
	return c.俱毗罗Kubera
}
    
func (c *俱毗罗KuberaCounty) BShitun石屯() feud.Barony {
	return c.石屯Shitun
}
    
func (c *俱毗罗KuberaCounty) BToqsu托克苏() feud.Barony {
	return c.托克苏Toqsu
}
    
func (c *俱毗罗KuberaCounty) BYangtaoshan杨桃山() feud.Barony {
	return c.杨桃山Yangtaoshan
}
    
var CKubera俱毗罗 KuberaCounty = &俱毗罗KuberaCounty{}

func init() {
	f := CKubera俱毗罗.(*俱毗罗KuberaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1528",
		Title:     "kubera",
		TitleName: "俱毗罗",
		TitleCode: "c_kubera",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿悉言Axiyan = BAxiyan阿悉言
	f.阿悉言Axiyan.SetParent(f)
	
	f.大州分场Dazhou_fenchang = BDazhou_fenchang大州分场
	f.大州分场Dazhou_fenchang.SetParent(f)
	
	f.葛仙庙Gexianmiao = BGexianmiao葛仙庙
	f.葛仙庙Gexianmiao.SetParent(f)
	
	f.俱毗罗Kubera = BKubera俱毗罗
	f.俱毗罗Kubera.SetParent(f)
	
	f.石屯Shitun = BShitun石屯
	f.石屯Shitun.SetParent(f)
	
	f.托克苏Toqsu = BToqsu托克苏
	f.托克苏Toqsu.SetParent(f)
	
	f.杨桃山Yangtaoshan = BYangtaoshan杨桃山
	f.杨桃山Yangtaoshan.SetParent(f)
	
}
