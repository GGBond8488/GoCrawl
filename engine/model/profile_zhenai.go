package model

import "encoding/json"

type PersonProfile struct {


	Name string
	Gender string

	Marriage string
	Age string
	Xinzuo string
	Height string
	Weight string
	WorkWhere string
	Income string
	Occupation string
	Education string

	Minzu string
	Jiguan string
	Figure string

	House string
	Car string
	Child string


}

type Models struct {
	Id string
	Url string
}

func FromJsonObject(o interface{})(PersonProfile,error)  {
	var person PersonProfile
	s,err := json.Marshal(0)
	if err != nil{
		return person,err
	}
	err = json.Unmarshal(s,&person)
	return person,err
}