package storage

type Storage interface {
	Save()
	PickRandom()
	Remove()
	IsExist()
}

type Page struckt{
	URL string
	UserName string
}