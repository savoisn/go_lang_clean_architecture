package infra

type FilePersistor struct {
	file_name string
	content   []string
}

func NewFilePersistor(file_name string) *FilePersistor {
	return &FilePersistor{file_name: file_name}
}

func (f *FilePersistor) Persist(something string) string {
	f.content = append(f.content, something)
	return something
}

func (f *FilePersistor) GetContent() []string {
	return f.content
}
