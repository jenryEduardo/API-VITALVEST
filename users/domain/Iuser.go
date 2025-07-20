
package domain


type Iuser interface{
	Save(user User)error
	Delete(id int)error
	Update(user User,id int)error
	Get()([]User,error)
}