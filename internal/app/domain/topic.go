package domain

type Topic interface {
	Get() []topic
	Find(id string) topic
	Create(topic topic) (topic, error)
	Update(id string, topic topic) error
	Delete(id string) error
}

type topic struct {
	Name        string
	Description string
	Subscribers []*subscription
}

func Get() []topic {
	return nil
}

func Find(id string) topic {
	return topic{}
}

func Create(t topic) (topic, error) {
	return t, nil
}

func Update(id string, topic topic) error {
	return nil
}

func Delete(id string) error {
	return nil
}
