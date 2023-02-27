package qohistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QohistanCounty interface {
    feud.County
    BBardaskan巴尔达斯坎() 	feud.Barony
    BBeyhaq拜哈格() 	feud.Barony
    BDasu达苏() 	feud.Barony
    BFariman法里曼() 	feud.Barony
    BGonabad戈纳巴德() 	feud.Barony
    BMahvalat马瓦拉特() 	feud.Barony
    BTorbatjam托尔巴特贾姆() 	feud.Barony
    BTorshiz图尔希兹() 	feud.Barony
}

type 忽希思丹QohistanCounty struct {
	feud.BaseCounty
	巴尔达斯坎Bardaskan 	feud.Barony
	拜哈格Beyhaq 	feud.Barony
	达苏Dasu 	feud.Barony
	法里曼Fariman 	feud.Barony
	戈纳巴德Gonabad 	feud.Barony
	马瓦拉特Mahvalat 	feud.Barony
	托尔巴特贾姆Torbatjam 	feud.Barony
	图尔希兹Torshiz 	feud.Barony
}

func (c *忽希思丹QohistanCounty) BBardaskan巴尔达斯坎() feud.Barony {
	return c.巴尔达斯坎Bardaskan
}
    
func (c *忽希思丹QohistanCounty) BBeyhaq拜哈格() feud.Barony {
	return c.拜哈格Beyhaq
}
    
func (c *忽希思丹QohistanCounty) BDasu达苏() feud.Barony {
	return c.达苏Dasu
}
    
func (c *忽希思丹QohistanCounty) BFariman法里曼() feud.Barony {
	return c.法里曼Fariman
}
    
func (c *忽希思丹QohistanCounty) BGonabad戈纳巴德() feud.Barony {
	return c.戈纳巴德Gonabad
}
    
func (c *忽希思丹QohistanCounty) BMahvalat马瓦拉特() feud.Barony {
	return c.马瓦拉特Mahvalat
}
    
func (c *忽希思丹QohistanCounty) BTorbatjam托尔巴特贾姆() feud.Barony {
	return c.托尔巴特贾姆Torbatjam
}
    
func (c *忽希思丹QohistanCounty) BTorshiz图尔希兹() feud.Barony {
	return c.图尔希兹Torshiz
}
    
var CQohistan忽希思丹 QohistanCounty = &忽希思丹QohistanCounty{}

func init() {
	f := CQohistan忽希思丹.(*忽希思丹QohistanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "635",
		Title:     "qohistan",
		TitleName: "忽希思丹",
		TitleCode: "c_qohistan",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔达斯坎Bardaskan = BBardaskan巴尔达斯坎
	f.巴尔达斯坎Bardaskan.SetParent(f)
	
	f.拜哈格Beyhaq = BBeyhaq拜哈格
	f.拜哈格Beyhaq.SetParent(f)
	
	f.达苏Dasu = BDasu达苏
	f.达苏Dasu.SetParent(f)
	
	f.法里曼Fariman = BFariman法里曼
	f.法里曼Fariman.SetParent(f)
	
	f.戈纳巴德Gonabad = BGonabad戈纳巴德
	f.戈纳巴德Gonabad.SetParent(f)
	
	f.马瓦拉特Mahvalat = BMahvalat马瓦拉特
	f.马瓦拉特Mahvalat.SetParent(f)
	
	f.托尔巴特贾姆Torbatjam = BTorbatjam托尔巴特贾姆
	f.托尔巴特贾姆Torbatjam.SetParent(f)
	
	f.图尔希兹Torshiz = BTorshiz图尔希兹
	f.图尔希兹Torshiz.SetParent(f)
	
}
