package main

type Partida struct {
	data          Data
	casa          Clube
	visitante     Clube
	golsCasa      int
	golsVisitante int
	publico       int
}

func (p *Partida) Jogar() {
	p.golsCasa = calcularGols(p.casa, true) // perfilTimeA + moralTimeA + moralJogadores + casa()
	p.golsVisitante = calcularGols(p.visitante, false)
}

func calcularGols(clube Clube, casa bool) int {
	var totalMoral int = 0
	var totalOthers int = 0

	// TODO: a habilidade do tecnico deve influenciar na partidade de forma forte

	for _, j := range clube.time.jogadores {
		totalMoral += j.moral
		totalOthers += j.perfil.aceleracao
		// TODO: calculate per category
	}

	total := totalMoral + totalOthers

	if casa {
		total = total + 100
	}

	if total > 700 {
		return 3 // TODO: randomizar
	}

	if total > 600 {
		return 2 // TODO: randomizar
	}

	if total > 500 {
		return 2 // TODO: randomizar
	}

	if total > 400 {
		return 1 // TODO: randomizar
	}

	if total > 300 {
		return 1 // TODO: randomizar
	}

	return 0
}
