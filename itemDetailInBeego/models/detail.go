package models

import "fmt"

//模仿返回数据库
type photo struct {
	Id        string
	ImagePath string
}

//返回数据 phoths部分

func GetPhotos() []photo {
	buffer := []photo{
		photo{
			Id:        "pic1",
			ImagePath: "img01.png",
		},
		photo{
			Id:        "pic2",
			ImagePath: "img02.png",
		},
		photo{
			Id:        "pic3",
			ImagePath: "img03.png",
		},
		photo{
			Id:        "pic4",
			ImagePath: "img04.png",
		},
		photo{
			Id:        "pic5",
			ImagePath: "img05.png",
		},
		photo{
			Id:        "pic6",
			ImagePath: "img06.png",
		},
	}
	return buffer
}

//返回数据 返回title
func GetTitle() string {
	return "北欧风格实木床"
}

//返回数据price
func GetPriceString() string {
	price := 000
	intPrice := price / 100
	decPrice := price % 100
	return fmt.Sprintf("%d.%02d", intPrice, decPrice)
}
