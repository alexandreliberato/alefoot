package main

import "fmt"

// TODO: temporada -> campeonatos x calendario

// um calendario tem um campeonato
// um campeonato de um calendário tem partidas entre vários times
// cada partida acontece em uma data do calendario do campeonato
// um campeonato tem varias datas para um calendario
// cada time de um campeonato têm vários jogadores
// cada jogador é de um time
// partidas entre times com jogadores em uma data de calendario

type Estadio struct {
	nome    string
	tamanho int
	preco   int
}

type Jogador struct {
	nome   string
	perfil PerfilTecnico
	// TODO: usar decimal
	valor   int     // em centavos
	moral   int     // 0-100
	posicao Posicao // 11 posicoes valiadas
}

type Posicao string

type Tecnico struct {
	nome       string
	habilidade int // 0-100
}

type PerfilTecnico struct {
	// 0-100 cada
	defesa       int
	penalti      int
	velocidade   int
	forca        int
	aceleracao   int
	drible       int
	desarme      int
	passe        int
	lancamento   int
	cruzamento   int
	chute        int
	cabeceio     int
	inteligencia int

	destro bool
}

type Calendario struct {
	campeonato Campeonato
	// iniciar no ano atual mas pode-se fazer logica mais complexa
	ano       int
	inicioMes int
	fimMes    int
}

type Campeonato struct {
	partidas []Partida
	clubes   []Clube // 20
}

type Data struct {
	ano int
	mes int
	dia int
}

func main() {

	jogador := Jogador{
		nome: "teste",
		perfil: PerfilTecnico{
			defesa:       50,
			penalti:      70,
			velocidade:   50,
			forca:        50,
			aceleracao:   50,
			drible:       50,
			desarme:      50,
			passe:        80,
			lancamento:   50,
			cruzamento:   50,
			chute:        50,
			cabeceio:     50,
			inteligencia: 50,
			destro:       true,
		},
	}

	jogadorB := Jogador{
		nome: "ruim",
		perfil: PerfilTecnico{
			defesa:       25,
			penalti:      25,
			velocidade:   25,
			forca:        25,
			aceleracao:   50,
			drible:       75,
			desarme:      15,
			passe:        25,
			lancamento:   45,
			cruzamento:   45,
			chute:        35,
			cabeceio:     55,
			inteligencia: 15,
			destro:       true,
		},
	}

	timeCorinthians, err := NewTime(
		Tecnico{
			nome:       "Wanderley Luxemburgo",
			habilidade: 80,
		},
		[]Jogador{
			jogador,
			jogador,
			jogador,
			jogador,
			jogador,
			jogador,
			jogador,
			jogador,
			jogador,
			jogador,
			jogador,
		},
	)
	if err != nil {
		panic(fmt.Errorf("criando time: %w", err))
	}

	clubeA := NewClube(
		"Corinthians",
		Estadio{
			nome:    "Arena Corinthians",
			tamanho: 50000,
			preco:   100,
		},
		timeCorinthians,
	)
	// pelo menos 22 jogadores cada time
	timePalmeiras, err := NewTime(
		Tecnico{
			nome:       "Lorem ipsum",
			habilidade: 60,
		},
		[]Jogador{
			jogadorB,
			jogadorB,
			jogadorB,
			jogadorB,
			jogadorB,
			jogadorB,
			jogadorB,
			jogadorB,
			jogadorB,
			jogadorB,
			jogadorB,
		},
	)
	if err != nil {
		panic(fmt.Errorf("criando time: %w", err))
	}

	clubeB := NewClube(
		"Palmeiras",
		Estadio{
			nome:    "Porqueiro",
			tamanho: 24000,
			preco:   200,
		},
		timePalmeiras,
	)

	// clubes := []Clube{
	// 	clubeA,
	// 	clubeB,
	// }

	// times nao podem ter os mesmos jogadores

	// campeonato := Campeonato{
	// 	clubes: clubes,
	// }

	// calendario := Calendario{
	// 	campeonato: campeonato,
	// 	ano:        2024,
	// 	inicioMes:  3,
	// 	fimMes:     12,
	// }

	partida := Partida{
		casa:      clubeA,
		visitante: clubeB,
		publico:   40000, // depende da moral do time
	}

	partida.Jogar()
}
