package domain

type PetRepository interface {
	Create(pet *Pet) error
	Update(pet *Pet) error
	Delete(id int64) error
	GetByID(id int64) (*Pet, error)
	FindByStatus(status []string) ([]Pet, error)
	FindByTags(tags []string) ([]Pet, error)
}

type UserRepository interface {
	Create(user *User) error
	CreateWithArray(users []User) error
	Update(username string, user *User) error
	Delete(username string) error
	GetByUsername(username string) (*User, error)
	Login(username, password string) (string, error)
	Logout() error
}

type StoreRepository interface {
	GetInventory() (map[string]int32, error)
	PlaceOrder(order *Order) error
	GetOrderByID(orderId int64) (*Order, error)
	DeleteOrder(orderId int64) error
}
