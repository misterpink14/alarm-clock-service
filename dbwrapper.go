package main

type DatabaseWrapper interface {
	Get()
	Create()
	Update()
	Delete()
}

type FakerWrapper struct {
}

func GetDatabaseWrapper() DatabaseWrapper {
	return FakerWrapper{}
}

func (fakerWrapper *FakerWrapper) Get(model Model) (Model, error) {
	return model
}

func (fakerWrapper *FakerWrapper) Create(model Model) (bool, error) {
	return true
}

func (fakerWrapper *FakerWrapper) Update(model Model) {
	return true
}

func (fakerWrapper *FakerWrapper) Delete(model Model) {
	return true
}
