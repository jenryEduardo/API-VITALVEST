
package domain


type Iuser interface{
	Save(user User)error
	Delete(id int)error
}