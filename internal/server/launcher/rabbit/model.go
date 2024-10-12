package rabbit

type consumer struct {
	topic string
}

type Model struct {
	consumers []consumer
}

func NewModel() Model {
	model := Model{}

	model.consumers = append(model.consumers, consumer{topic: "yellow"})

	return model
}
