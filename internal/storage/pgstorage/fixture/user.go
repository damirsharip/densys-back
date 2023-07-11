package fixture

import (
	"time"

	"github.com/khanfromasia/densys/admin/internal/entity"
)

type UserBuilder struct {
	instance *entity.User
}

func User() *UserBuilder {
	return &UserBuilder{
		instance: &entity.User{},
	}
}

func (b *UserBuilder) ID(v string) *UserBuilder {
	b.instance.ID = v
	return b
}

func (b *UserBuilder) FirstName(v string) *UserBuilder {
	b.instance.FirstName = v
	return b
}

func (b *UserBuilder) LastName(v string) *UserBuilder {
	b.instance.LastName = v
	return b
}

func (b *UserBuilder) MiddleName(v string) *UserBuilder {
	b.instance.MiddleName = v
	return b
}

func (b *UserBuilder) Address(v string) *UserBuilder {
	b.instance.Address = v
	return b
}

func (b *UserBuilder) Email(v string) *UserBuilder {
	b.instance.Email = v
	return b
}

func (b *UserBuilder) Password(v string) *UserBuilder {
	b.instance.Password = v
	return b
}

func (b *UserBuilder) BirthDate(v time.Time) *UserBuilder {
	b.instance.BirthDate = v
	return b
}

func (b *UserBuilder) Role(v string) *UserBuilder {
	b.instance.Role = v
	return b
}

func (b *UserBuilder) PhoneNumber(v string) *UserBuilder {
	b.instance.PhoneNumber = v
	return b
}

func (b *UserBuilder) IIN(v string) *UserBuilder {
	b.instance.PhoneNumber = v
	return b
}

func (b *UserBuilder) GovernmentID(v string) *UserBuilder {
	b.instance.PhoneNumber = v
	return b
}

func (b *UserBuilder) CreatedAt(v time.Time) *UserBuilder {
	b.instance.CreatedAt = v
	return b
}

func (b *UserBuilder) P() *entity.User {
	return b.instance
}

func (b *UserBuilder) V() entity.User {
	return *b.instance
}
