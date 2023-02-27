package mahoyadapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MahoyadapuramCounty interface {
    feud.County
    BKadungallur迦顿迦卢尔() 	feud.Barony
    BKaladi伽罗迪() 	feud.Barony
    BKaviyur伽毗由() 	feud.Barony
    BKunjakari贡奢迦利() 	feud.Barony
    BMahoyadapuram摩呼陀耶补罗() 	feud.Barony
    BTripunittura帝利补尼菟罗() 	feud.Barony
    BUdagai尤陀伽() 	feud.Barony
}

type 摩呼陀耶补罗MahoyadapuramCounty struct {
	feud.BaseCounty
	迦顿迦卢尔Kadungallur 	feud.Barony
	伽罗迪Kaladi 	feud.Barony
	伽毗由Kaviyur 	feud.Barony
	贡奢迦利Kunjakari 	feud.Barony
	摩呼陀耶补罗Mahoyadapuram 	feud.Barony
	帝利补尼菟罗Tripunittura 	feud.Barony
	尤陀伽Udagai 	feud.Barony
}

func (c *摩呼陀耶补罗MahoyadapuramCounty) BKadungallur迦顿迦卢尔() feud.Barony {
	return c.迦顿迦卢尔Kadungallur
}
    
func (c *摩呼陀耶补罗MahoyadapuramCounty) BKaladi伽罗迪() feud.Barony {
	return c.伽罗迪Kaladi
}
    
func (c *摩呼陀耶补罗MahoyadapuramCounty) BKaviyur伽毗由() feud.Barony {
	return c.伽毗由Kaviyur
}
    
func (c *摩呼陀耶补罗MahoyadapuramCounty) BKunjakari贡奢迦利() feud.Barony {
	return c.贡奢迦利Kunjakari
}
    
func (c *摩呼陀耶补罗MahoyadapuramCounty) BMahoyadapuram摩呼陀耶补罗() feud.Barony {
	return c.摩呼陀耶补罗Mahoyadapuram
}
    
func (c *摩呼陀耶补罗MahoyadapuramCounty) BTripunittura帝利补尼菟罗() feud.Barony {
	return c.帝利补尼菟罗Tripunittura
}
    
func (c *摩呼陀耶补罗MahoyadapuramCounty) BUdagai尤陀伽() feud.Barony {
	return c.尤陀伽Udagai
}
    
var CMahoyadapuram摩呼陀耶补罗 MahoyadapuramCounty = &摩呼陀耶补罗MahoyadapuramCounty{}

func init() {
	f := CMahoyadapuram摩呼陀耶补罗.(*摩呼陀耶补罗MahoyadapuramCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1114",
		Title:     "mahoyadapuram",
		TitleName: "摩呼陀耶补罗",
		TitleCode: "c_mahoyadapuram",
		Baronies:  map[string]feud.Barony{},
	}

	f.迦顿迦卢尔Kadungallur = BKadungallur迦顿迦卢尔
	f.迦顿迦卢尔Kadungallur.SetParent(f)
	
	f.伽罗迪Kaladi = BKaladi伽罗迪
	f.伽罗迪Kaladi.SetParent(f)
	
	f.伽毗由Kaviyur = BKaviyur伽毗由
	f.伽毗由Kaviyur.SetParent(f)
	
	f.贡奢迦利Kunjakari = BKunjakari贡奢迦利
	f.贡奢迦利Kunjakari.SetParent(f)
	
	f.摩呼陀耶补罗Mahoyadapuram = BMahoyadapuram摩呼陀耶补罗
	f.摩呼陀耶补罗Mahoyadapuram.SetParent(f)
	
	f.帝利补尼菟罗Tripunittura = BTripunittura帝利补尼菟罗
	f.帝利补尼菟罗Tripunittura.SetParent(f)
	
	f.尤陀伽Udagai = BUdagai尤陀伽
	f.尤陀伽Udagai.SetParent(f)
	
}
