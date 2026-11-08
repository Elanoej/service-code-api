package main

type Servico struct {
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}

func MockServicos() []Servico {
	return []Servico{
		{Codigo: "010101", Descricao: "Análise e desenvolvimento de sistemas."},
		{Codigo: "010201", Descricao: "Programação."},
		{Codigo: "010301", Descricao: "Processamento de dados, textos, imagens, vídeos, páginas eletrônicas, aplicativos e sistemas de informação, entre outros formatos, e congêneres."},
		{Codigo: "010302", Descricao: "Armazenamento ou hospedagem de dados, textos, imagens, vídeos, páginas eletrônicas, aplicativos e sistemas de informação, entre outros formatos, e congêneres."},
		{Codigo: "010401", Descricao: "Elaboração de programas de computadores, inclusive de jogos eletrônicos, independentemente da arquitetura construtiva da máquina em que o programa será executado, incluindo tablets, smartphones e congêneres."},
	}
}
