package uuid

import uuid "github.com/satori/go.uuid"

type UUID [16]byte

func (u UUID) String() string {
	impl := uuid.UUID(u)
	return impl.String()
}

func FromString(input string) (u UUID, err error) {
	impl, err := uuid.FromString(input)
	if err != nil {
		return u, err
	}
	u = UUID(impl)
	return
}

func Generate() UUID {
	impl := uuid.NewV4()
	return UUID(impl)
}
