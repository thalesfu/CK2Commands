package artux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArtuxCounty interface {
	feud.County
	BAtush阿图什() feud.Barony
	BChangjiang昌疆() feud.Barony
	BCuocaogou错草沟() feud.Barony
	BDalihe答丽河() feud.Barony
	BHugou湖沟() feud.Barony
	BPoskam波斯喀木() feud.Barony
	BYanzi岩姿() feud.Barony
}

type 阿图什ArtuxCounty struct {
	feud.BaseCounty
	阿图什Atush     feud.Barony
	昌疆Changjiang feud.Barony
	错草沟Cuocaogou feud.Barony
	答丽河Dalihe    feud.Barony
	湖沟Hugou      feud.Barony
	波斯喀木Poskam   feud.Barony
	岩姿Yanzi      feud.Barony
}

func (c *阿图什ArtuxCounty) BAtush阿图什() feud.Barony {
	return c.阿图什Atush
}

func (c *阿图什ArtuxCounty) BChangjiang昌疆() feud.Barony {
	return c.昌疆Changjiang
}

func (c *阿图什ArtuxCounty) BCuocaogou错草沟() feud.Barony {
	return c.错草沟Cuocaogou
}

func (c *阿图什ArtuxCounty) BDalihe答丽河() feud.Barony {
	return c.答丽河Dalihe
}

func (c *阿图什ArtuxCounty) BHugou湖沟() feud.Barony {
	return c.湖沟Hugou
}

func (c *阿图什ArtuxCounty) BPoskam波斯喀木() feud.Barony {
	return c.波斯喀木Poskam
}

func (c *阿图什ArtuxCounty) BYanzi岩姿() feud.Barony {
	return c.岩姿Yanzi
}

var CArtux阿图什 ArtuxCounty = &阿图什ArtuxCounty{}

func init() {
	f := CArtux阿图什.(*阿图什ArtuxCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1526",
		Title:     "artux",
		TitleName: "阿图什",
		TitleCode: "c_artux",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿图什Atush = BAtush阿图什
	f.阿图什Atush.SetParent(f)

	f.昌疆Changjiang = BChangjiang昌疆
	f.昌疆Changjiang.SetParent(f)

	f.错草沟Cuocaogou = BCuocaogou错草沟
	f.错草沟Cuocaogou.SetParent(f)

	f.答丽河Dalihe = BDalihe答丽河
	f.答丽河Dalihe.SetParent(f)

	f.湖沟Hugou = BHugou湖沟
	f.湖沟Hugou.SetParent(f)

	f.波斯喀木Poskam = BPoskam波斯喀木
	f.波斯喀木Poskam.SetParent(f)

	f.岩姿Yanzi = BYanzi岩姿
	f.岩姿Yanzi.SetParent(f)

}
